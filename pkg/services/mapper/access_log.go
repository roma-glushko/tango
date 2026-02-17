package mapper

import (
	"fmt"
	"strconv"
	"strings"
	"tango/pkg/entity"
	"time"
)

var timeFormat = "02/Jan/2006:15:04:05 -0700"

// ParseApacheCombined parses an Apache/NGINX Combined Log Format line.
// Format: ip_list - identity [time] "method uri protocol" code size "referer" "user_agent"
func ParseApacheCombined(line string) (entity.AccessLogRecord, error) {
	line = strings.TrimSpace(line)
	if len(line) == 0 {
		return entity.AccessLogRecord{}, fmt.Errorf("empty log line")
	}

	// 1. Find " - [" to split IP list from identity and time fields.
	// Format: "ip1 ip2, ip3 - [time] ..."
	// The " - [" separates: ip_list (space) identity (space) - (space) [
	// But the identity is always "-", so we match on " - [" directly.
	idx := strings.Index(line, " - [")
	if idx < 0 {
		return entity.AccessLogRecord{}, fmt.Errorf("missing identity field in: %s", line)
	}

	ipListRaw := line[:idx]
	rest := line[idx+4:] // skip " - ["

	// 2. Find "] " to extract time
	idx = strings.Index(rest, "] \"")
	if idx < 0 {
		return entity.AccessLogRecord{}, fmt.Errorf("missing time closing bracket in: %s", line)
	}
	timeStr := rest[:idx]
	rest = rest[idx+3:] // skip "] "

	// 3. Find closing quote of request line
	idx = strings.Index(rest, "\" ")
	if idx < 0 {
		return entity.AccessLogRecord{}, fmt.Errorf("missing request closing quote in: %s", line)
	}
	requestLine := rest[:idx]
	rest = rest[idx+2:] // skip "\" "

	// 4. Split request line: "GET /path HTTP/1.1"
	var method, uri, protocol string
	parts := strings.SplitN(requestLine, " ", 3)
	method = parts[0]
	if len(parts) > 1 {
		uri = parts[1]
	}
	if len(parts) > 2 {
		protocol = parts[2]
	}

	// 5. Extract response code (next token before space)
	idx = strings.IndexByte(rest, ' ')
	if idx < 0 {
		return entity.AccessLogRecord{}, fmt.Errorf("missing response code in: %s", line)
	}
	responseCodeStr := rest[:idx]
	rest = rest[idx+1:]

	// 6. Extract response size (next token before space or end)
	idx = strings.IndexByte(rest, ' ')
	var responseSizeStr string
	if idx < 0 {
		// Common log format (no referer/ua) - rest is just the size
		responseSizeStr = rest
		rest = ""
	} else {
		responseSizeStr = rest[:idx]
		rest = rest[idx+1:]
	}

	// 7. Extract referer (quoted) and user agent (quoted)
	referer, userAgent := parseQuotedFields(rest)

	// Parse fields into typed values
	ipStr := cleanIPList(ipListRaw)

	parsedTime, err := time.Parse(timeFormat, timeStr)
	if err != nil {
		return entity.AccessLogRecord{}, fmt.Errorf("failed to parse time %q: %w", timeStr, err)
	}

	responseCode, err := parseUintOrDash(responseCodeStr)
	if err != nil {
		return entity.AccessLogRecord{}, fmt.Errorf("failed to parse response code %q: %w", responseCodeStr, err)
	}

	responseSize, err := parseUintOrDash(responseSizeStr)
	if err != nil {
		return entity.AccessLogRecord{}, fmt.Errorf("failed to parse response size %q: %w", responseSizeStr, err)
	}

	return entity.AccessLogRecord{
		IP:            ipStr,
		URI:           uri,
		Time:          parsedTime,
		RequestMethod: method,
		Protocol:      protocol,
		ResponseCode:  responseCode,
		ResponseSize:  responseSize,
		RefererURL:    referer,
		UserAgent:     userAgent,
	}, nil
}

// ParseApacheCommon parses an Apache Common Log Format line (no referer/user_agent).
// Format: ip_list identity - [time] "method uri protocol" code size
func ParseApacheCommon(line string) (entity.AccessLogRecord, error) {
	line = strings.TrimSpace(line)
	if len(line) == 0 {
		return entity.AccessLogRecord{}, fmt.Errorf("empty log line")
	}

	idx := strings.Index(line, " - [")
	if idx < 0 {
		return entity.AccessLogRecord{}, fmt.Errorf("missing identity field in: %s", line)
	}

	ipListRaw := line[:idx]
	rest := line[idx+4:]

	idx = strings.Index(rest, "] \"")
	if idx < 0 {
		return entity.AccessLogRecord{}, fmt.Errorf("missing time closing bracket in: %s", line)
	}
	timeStr := rest[:idx]
	rest = rest[idx+3:]

	idx = strings.Index(rest, "\" ")
	if idx < 0 {
		return entity.AccessLogRecord{}, fmt.Errorf("missing request closing quote in: %s", line)
	}
	requestLine := rest[:idx]
	rest = rest[idx+2:]

	var method, uri, protocol string
	parts := strings.SplitN(requestLine, " ", 3)
	method = parts[0]
	if len(parts) > 1 {
		uri = parts[1]
	}
	if len(parts) > 2 {
		protocol = parts[2]
	}

	idx = strings.IndexByte(rest, ' ')
	if idx < 0 {
		return entity.AccessLogRecord{}, fmt.Errorf("missing response code in: %s", line)
	}
	responseCodeStr := rest[:idx]
	responseSizeStr := rest[idx+1:]

	// Trim any trailing whitespace from response size
	responseSizeStr = strings.TrimSpace(responseSizeStr)

	ipStr := cleanIPList(ipListRaw)

	parsedTime, err := time.Parse(timeFormat, timeStr)
	if err != nil {
		return entity.AccessLogRecord{}, fmt.Errorf("failed to parse time %q: %w", timeStr, err)
	}

	responseCode, err := parseUintOrDash(responseCodeStr)
	if err != nil {
		return entity.AccessLogRecord{}, fmt.Errorf("failed to parse response code %q: %w", responseCodeStr, err)
	}

	responseSize, err := parseUintOrDash(responseSizeStr)
	if err != nil {
		return entity.AccessLogRecord{}, fmt.Errorf("failed to parse response size %q: %w", responseSizeStr, err)
	}

	return entity.AccessLogRecord{
		IP:            ipStr,
		URI:           uri,
		Time:          parsedTime,
		RequestMethod: method,
		Protocol:      protocol,
		ResponseCode:  responseCode,
		ResponseSize:  responseSize,
	}, nil
}

// cleanIPList normalizes an IP list string, removing dashes and extra separators.
func cleanIPList(raw string) string {
	normalized := strings.ReplaceAll(raw, ", ", " ")

	// Fast path: single IP with no dashes
	if !strings.ContainsAny(normalized, " -") {
		return normalized
	}

	parts := strings.Split(normalized, " ")
	var b strings.Builder
	for _, p := range parts {
		if p != "-" && p != "" {
			if b.Len() > 0 {
				b.WriteByte(' ')
			}
			b.WriteString(p)
		}
	}
	return b.String()
}

// parseQuotedFields extracts referer and user agent from: "referer" "user_agent"
func parseQuotedFields(s string) (referer, userAgent string) {
	if len(s) == 0 {
		return "", ""
	}

	// Expect: "referer" "user_agent"
	// Find first quoted string
	if s[0] == '"' {
		idx := strings.Index(s[1:], "\"")
		if idx >= 0 {
			referer = s[1 : idx+1]
			rest := s[idx+2:] // skip closing quote

			// Find second quoted string
			start := strings.IndexByte(rest, '"')
			if start >= 0 {
				end := strings.LastIndexByte(rest, '"')
				if end > start {
					userAgent = rest[start+1 : end]
				}
			}
		}
	}

	return referer, userAgent
}

// parseUintOrDash parses a uint64 from a string, treating "-" as 0.
func parseUintOrDash(s string) (uint64, error) {
	if s == "-" {
		return 0, nil
	}

	return strconv.ParseUint(s, 10, 64)
}
