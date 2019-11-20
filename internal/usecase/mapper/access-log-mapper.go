package mapper

import (
	"regexp"
	"strconv"
	"strings"
	"tango/internal/domain/entity"
	"time"
)

var timeFormat = "02/Jan/2006:15:04:05 -0700"
var combinedLogFormat = `^(?P<ip_list>[\S, ]+) (\-) \[(?P<time>[\w:/]+\s[+\-]\d{4})\] "(?P<request_method>\S+)\s?(?P<uri>\S+)?\s?(?P<protocol>\S+)?" (?P<response_code>\d{3}|-) (?P<response_size>\d+|-)\s?"?(?P<referer_url>[^"]*)"?\s?"?(?P<user_agent>[^"]*)?"?$`

func filter(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func findNamedMatches(regex *regexp.Regexp, str string) map[string]string {
	match := regex.FindStringSubmatch(str)

	results := map[string]string{}

	for i, name := range match {
		results[regex.SubexpNames()[i]] = name
	}

	return results
}

// Map access logs line to AccessLogRecord type
func MapAccessLogRecord(accessLogRecord string) entity.AccessLogRecord {
	// todo: move compiling from the map method, we need it once and then use compiled pattern
	accessLogParser, _ := regexp.Compile(combinedLogFormat)

	accessRecordInformation := findNamedMatches(accessLogParser, strings.TrimSpace(accessLogRecord))

	ipList := filter(
		strings.Split(
			strings.ReplaceAll(accessRecordInformation["ip_list"], ", ", " "),
			" ",
		),
		"-",
	)

	time, _ := time.Parse(timeFormat, accessRecordInformation["time"])

	responseCode, _ := strconv.ParseUint(accessRecordInformation["response_code"], 10, 64)
	responseSize, _ := strconv.ParseUint(accessRecordInformation["response_size"], 10, 64)

	return entity.AccessLogRecord{
		IP:            ipList,
		URI:           accessRecordInformation["uri"],
		Time:          time,
		RequestMethod: accessRecordInformation["request_method"],
		Protocol:      accessRecordInformation["protocol"],
		ResponseCode:  responseCode,
		ResponseSize:  responseSize,
		RefererURL:    accessRecordInformation["referer_url"],
		UserAgent:     accessRecordInformation["user_agent"],
	}
}
