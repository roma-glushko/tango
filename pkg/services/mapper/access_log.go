package mapper

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"tango/pkg/entity"
	"time"
)

var timeFormat = "02/Jan/2006:15:04:05 -0700"
var combinedLogFormat = `^(?P<ip_list>[\S, ]+) (\-) \[(?P<time>[\w:/]+\s[+\-]\d{4})\] "(?P<request_method>\S+)\s?(?P<uri>\S+)?\s?(?P<protocol>\S+)?" (?P<response_code>\d{3}|-) (?P<response_size>\d+|-)\s?"?(?P<referer_url>[^"]*)"?\s?"?(?P<user_agent>[^"]*)?"?$`
var accessLogParser = regexp.MustCompile(combinedLogFormat)

func filter(s []string, r string) []string {
	result := make([]string, 0, len(s))
	for _, v := range s {
		if v != r {
			result = append(result, v)
		}
	}
	return result
}

func findNamedMatches(regex *regexp.Regexp, str string) map[string]string {
	match := regex.FindStringSubmatch(str)

	results := map[string]string{}

	for i, name := range match {
		results[regex.SubexpNames()[i]] = name
	}

	return results
}

// MapAccessLogRecord maps access log line to AccessLogRecord type
func MapAccessLogRecord(accessLogRecord string) (entity.AccessLogRecord, error) {
	accessRecordInformation := findNamedMatches(accessLogParser, strings.TrimSpace(accessLogRecord))

	if len(accessRecordInformation) == 0 {
		return entity.AccessLogRecord{}, fmt.Errorf("failed to parse access log line: %s", accessLogRecord)
	}

	ipList := filter(
		strings.Split(
			strings.ReplaceAll(accessRecordInformation["ip_list"], ", ", " "),
			" ",
		),
		"-",
	)

	parsedTime, err := time.Parse(timeFormat, accessRecordInformation["time"])
	if err != nil {
		return entity.AccessLogRecord{}, fmt.Errorf("failed to parse time %q: %w", accessRecordInformation["time"], err)
	}

	var responseCode uint64
	if accessRecordInformation["response_code"] != "-" {
		responseCode, err = strconv.ParseUint(accessRecordInformation["response_code"], 10, 64)
		if err != nil {
			return entity.AccessLogRecord{}, fmt.Errorf("failed to parse response code %q: %w", accessRecordInformation["response_code"], err)
		}
	}

	var responseSize uint64
	if accessRecordInformation["response_size"] != "-" {
		responseSize, err = strconv.ParseUint(accessRecordInformation["response_size"], 10, 64)
		if err != nil {
			return entity.AccessLogRecord{}, fmt.Errorf("failed to parse response size %q: %w", accessRecordInformation["response_size"], err)
		}
	}

	return entity.AccessLogRecord{
		IP:            ipList,
		URI:           accessRecordInformation["uri"],
		Time:          parsedTime,
		RequestMethod: accessRecordInformation["request_method"],
		Protocol:      accessRecordInformation["protocol"],
		ResponseCode:  responseCode,
		ResponseSize:  responseSize,
		RefererURL:    accessRecordInformation["referer_url"],
		UserAgent:     accessRecordInformation["user_agent"],
	}, nil
}
