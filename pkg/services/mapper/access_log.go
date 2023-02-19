package mapper

import (
	"regexp"
	"strconv"
	"strings"
	"sync"
	"tango/pkg/entity"
	"time"
)

var timeFormat = "02/Jan/2006:15:04:05 -0700"
var combinedLogFormat = `^(?P<ip_list>[\S, ]+) (\-) \[(?P<time>[\w:/]+\s[+\-]\d{4})\] "(?P<request_method>\S+)\s?(?P<uri>\S+)?\s?(?P<protocol>\S+)?" (?P<response_code>\d{3}|-) (?P<response_size>\d+|-)\s?"?(?P<referer_url>[^"]*)"?\s?"?(?P<user_agent>[^"]*)?"?$`

type AccessLogMapper struct {
	pool   sync.Pool
	parser *regexp.Regexp
}

func NewAccessLogMapper() *AccessLogMapper {
	accessLogParser, _ := regexp.Compile(combinedLogFormat)

	return &AccessLogMapper{
		pool: sync.Pool{
			New: func() interface{} {
				return entity.AccessLogRecord{}
			},
		},
		parser: accessLogParser,
	}
}

// Map access logs line to AccessLogRecord type
func (m *AccessLogMapper) Map(recordLine string) entity.AccessLogRecord {
	accessRecordInformation := m.parseLine(m.parser, strings.TrimSpace(recordLine))

	ipList := m.filter(
		strings.Split(
			strings.ReplaceAll(accessRecordInformation["ip_list"], ", ", " "),
			" ",
		),
		"-",
	)

	createdAt, _ := time.Parse(timeFormat, accessRecordInformation["time"])

	responseCode, _ := strconv.ParseUint(accessRecordInformation["response_code"], 10, 64)
	responseSize, _ := strconv.ParseUint(accessRecordInformation["response_size"], 10, 64)

	accessLogRecord := m.pool.Get().(entity.AccessLogRecord)

	accessLogRecord.IP = ipList
	accessLogRecord.URI = accessRecordInformation["uri"]
	accessLogRecord.Time = createdAt
	accessLogRecord.RequestMethod = accessRecordInformation["request_method"]
	accessLogRecord.Protocol = accessRecordInformation["protocol"]
	accessLogRecord.ResponseCode = responseCode
	accessLogRecord.ResponseSize = responseSize
	accessLogRecord.RefererURL = accessRecordInformation["referer_url"]
	accessLogRecord.UserAgent = accessRecordInformation["user_agent"]

	return accessLogRecord
}

func (m *AccessLogMapper) parseLine(regex *regexp.Regexp, str string) map[string]string {
	match := regex.FindStringSubmatch(str)

	results := map[string]string{}

	for i, name := range match {
		results[regex.SubexpNames()[i]] = name
	}

	return results
}

func (m *AccessLogMapper) filter(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func (m *AccessLogMapper) Recycle(logRecord entity.AccessLogRecord) {
	m.pool.Put(logRecord)
}
