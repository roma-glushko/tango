package mapper

import (
	"bytes"
	"fmt"
	"strconv"
	"tango/pkg/entity"
	"time"
)

var timeFormat = "02/Jan/2006:15:04:05 -0700"

var (
	sepIdentity = []byte(" - [")
	sepTimEnd   = []byte("] \"")
	sepReqEnd   = []byte("\" ")
)

// ParseApacheCombined parses an Apache/NGINX Combined Log Format line.
// Format: ip_list - identity [time] "method uri protocol" code size "referer" "user_agent"
func ParseApacheCombined(line []byte) (entity.AccessLogRecord, error) {
	line = bytes.TrimSpace(line)
	if len(line) == 0 {
		return entity.AccessLogRecord{}, fmt.Errorf("empty log line")
	}

	// 1. Find " - [" to split IP list from identity and time fields.
	idx := bytes.Index(line, sepIdentity)
	if idx < 0 {
		return entity.AccessLogRecord{}, fmt.Errorf("missing identity field in: %s", line)
	}

	ipListRaw := line[:idx]
	rest := line[idx+4:] // skip " - ["

	// 2. Find "] " to extract time
	idx = bytes.Index(rest, sepTimEnd)
	if idx < 0 {
		return entity.AccessLogRecord{}, fmt.Errorf("missing time closing bracket in: %s", line)
	}
	timeSlice := rest[:idx]
	rest = rest[idx+3:] // skip "] "

	// 3. Find closing quote of request line
	idx = bytes.Index(rest, sepReqEnd)
	if idx < 0 {
		return entity.AccessLogRecord{}, fmt.Errorf("missing request closing quote in: %s", line)
	}
	requestLine := rest[:idx]
	rest = rest[idx+2:] // skip "\" "

	// 4. Split request line: "GET /path HTTP/1.1"
	var method, uri, protocol string
	if sp1 := bytes.IndexByte(requestLine, ' '); sp1 < 0 {
		method = string(requestLine)
	} else {
		method = string(requestLine[:sp1])
		rem := requestLine[sp1+1:]
		if sp2 := bytes.IndexByte(rem, ' '); sp2 < 0 {
			uri = string(rem)
		} else {
			uri = string(rem[:sp2])
			protocol = string(rem[sp2+1:])
		}
	}

	// 5. Extract response code (next token before space)
	idx = bytes.IndexByte(rest, ' ')
	if idx < 0 {
		return entity.AccessLogRecord{}, fmt.Errorf("missing response code in: %s", line)
	}
	responseCodeSlice := rest[:idx]
	rest = rest[idx+1:]

	// 6. Extract response size (next token before space or end)
	idx = bytes.IndexByte(rest, ' ')
	var responseSizeSlice []byte
	var trailing []byte
	if idx < 0 {
		responseSizeSlice = rest
	} else {
		responseSizeSlice = rest[:idx]
		trailing = rest[idx+1:]
	}

	// 7. Extract referer (quoted) and user agent (quoted)
	referer, userAgent := parseQuotedFields(trailing)

	// Parse fields into typed values
	ipStr := cleanIPList(ipListRaw)

	parsedTime, err := time.Parse(timeFormat, string(timeSlice))
	if err != nil {
		return entity.AccessLogRecord{}, fmt.Errorf("failed to parse time %q: %w", timeSlice, err)
	}

	responseCode, err := parseUintOrDash(responseCodeSlice)
	if err != nil {
		return entity.AccessLogRecord{}, fmt.Errorf("failed to parse response code %q: %w", responseCodeSlice, err)
	}

	responseSize, err := parseUintOrDash(responseSizeSlice)
	if err != nil {
		return entity.AccessLogRecord{}, fmt.Errorf("failed to parse response size %q: %w", responseSizeSlice, err)
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
func ParseApacheCommon(line []byte) (entity.AccessLogRecord, error) {
	line = bytes.TrimSpace(line)
	if len(line) == 0 {
		return entity.AccessLogRecord{}, fmt.Errorf("empty log line")
	}

	idx := bytes.Index(line, sepIdentity)
	if idx < 0 {
		return entity.AccessLogRecord{}, fmt.Errorf("missing identity field in: %s", line)
	}

	ipListRaw := line[:idx]
	rest := line[idx+4:]

	idx = bytes.Index(rest, sepTimEnd)
	if idx < 0 {
		return entity.AccessLogRecord{}, fmt.Errorf("missing time closing bracket in: %s", line)
	}
	timeSlice := rest[:idx]
	rest = rest[idx+3:]

	idx = bytes.Index(rest, sepReqEnd)
	if idx < 0 {
		return entity.AccessLogRecord{}, fmt.Errorf("missing request closing quote in: %s", line)
	}
	requestLine := rest[:idx]
	rest = rest[idx+2:]

	var method, uri, protocol string
	if sp1 := bytes.IndexByte(requestLine, ' '); sp1 < 0 {
		method = string(requestLine)
	} else {
		method = string(requestLine[:sp1])
		rem := requestLine[sp1+1:]
		if sp2 := bytes.IndexByte(rem, ' '); sp2 < 0 {
			uri = string(rem)
		} else {
			uri = string(rem[:sp2])
			protocol = string(rem[sp2+1:])
		}
	}

	idx = bytes.IndexByte(rest, ' ')
	if idx < 0 {
		return entity.AccessLogRecord{}, fmt.Errorf("missing response code in: %s", line)
	}
	responseCodeSlice := rest[:idx]
	responseSizeSlice := bytes.TrimSpace(rest[idx+1:])

	ipStr := cleanIPList(ipListRaw)

	parsedTime, err := time.Parse(timeFormat, string(timeSlice))
	if err != nil {
		return entity.AccessLogRecord{}, fmt.Errorf("failed to parse time %q: %w", timeSlice, err)
	}

	responseCode, err := parseUintOrDash(responseCodeSlice)
	if err != nil {
		return entity.AccessLogRecord{}, fmt.Errorf("failed to parse response code %q: %w", responseCodeSlice, err)
	}

	responseSize, err := parseUintOrDash(responseSizeSlice)
	if err != nil {
		return entity.AccessLogRecord{}, fmt.Errorf("failed to parse response size %q: %w", responseSizeSlice, err)
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

// cleanIPList normalizes an IP list from []byte, removing dashes and extra separators.
func cleanIPList(raw []byte) string {
	normalized := bytes.ReplaceAll(raw, []byte(", "), []byte(" "))

	// Fast path: single IP with no dashes
	if !bytes.ContainsAny(normalized, " -") {
		return string(normalized)
	}

	parts := bytes.Split(normalized, []byte(" "))
	var b bytes.Buffer
	for _, p := range parts {
		if len(p) > 0 && !(len(p) == 1 && p[0] == '-') {
			if b.Len() > 0 {
				b.WriteByte(' ')
			}
			b.Write(p)
		}
	}
	return b.String()
}

// parseQuotedFields extracts referer and user agent from: "referer" "user_agent"
func parseQuotedFields(s []byte) (referer, userAgent string) {
	if len(s) == 0 {
		return "", ""
	}

	if s[0] == '"' {
		idx := bytes.IndexByte(s[1:], '"')
		if idx >= 0 {
			referer = string(s[1 : idx+1])
			rest := s[idx+2:]

			start := bytes.IndexByte(rest, '"')
			if start >= 0 {
				end := bytes.LastIndexByte(rest, '"')
				if end > start {
					userAgent = string(rest[start+1 : end])
				}
			}
		}
	}

	return referer, userAgent
}

// parseUintOrDash parses a uint64 from a byte slice, treating "-" as 0.
func parseUintOrDash(s []byte) (uint64, error) {
	if len(s) == 1 && s[0] == '-' {
		return 0, nil
	}

	return strconv.ParseUint(string(s), 10, 64)
}
