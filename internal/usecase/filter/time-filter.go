package filter

import (
	"tango/internal/domain/entity"
	"time"
)

var timeFromFilter = "2019-07-08 12:00:00 -0400" // nil
var timeToFilter = "2019-07-08 13:00:00 -0400"   // nil

var europeFormat = "2006-01-02 15:04:05 -0700"

//
type TimeFilter struct {
}

//
func NewTimeFilter() *TimeFilter {
	return &TimeFilter{}
}

//
func (f *TimeFilter) Filter(accessLogRecord entity.AccessLogRecord) bool {
	if timeFromFilter == "" && timeToFilter == "" {
		return false
	}

	recordTime := accessLogRecord.Time

	timeStart, _ := time.Parse(europeFormat, timeFromFilter)
	timeEnd, _ := time.Parse(europeFormat, timeToFilter)

	if recordTime.After(timeStart) && recordTime.Before(timeEnd) {
		return false
	}

	return true
}
