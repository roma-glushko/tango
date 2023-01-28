package entity

import (
	"time"
)

// AccessLogRecord Type of AccessLog record
type AccessLogRecord struct {
	IP            []string
	URI           string
	Time          time.Time
	RequestMethod string
	Protocol      string
	ResponseCode  uint64
	ResponseSize  uint64
	UserAgent     string
	RefererURL    string
}
