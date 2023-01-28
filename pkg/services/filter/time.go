package filter

import (
	"tango/pkg/domain/entity"
	"tango/pkg/services/config"
	"time"
)

var EuropeFormat = "2006-01-02 15:04:05 -0700"

//
type TimeFilter struct {
	timeStart time.Time
	timeEnd   time.Time
}

//
func NewTimeFilter(filterConfig config.FilterConfig) *TimeFilter {
	timeStart := time.Time{}
	timeEnd := time.Time{}

	timeFrames := filterConfig.KeepTimeFrames

	// doesn't support multiple time frames by this moment
	// needs to validate time frames: timeStart should be less than timeEnd
	if len(timeFrames) == 2 {
		timeStart, _ = time.Parse(EuropeFormat, timeFrames[0])
		timeEnd, _ = time.Parse(EuropeFormat, timeFrames[1])
	}

	return &TimeFilter{
		timeStart: timeStart,
		timeEnd:   timeEnd,
	}
}

//
func (f *TimeFilter) Filter(accessLogRecord entity.AccessLogRecord) bool {
	if f.timeStart.IsZero() && f.timeEnd.IsZero() {
		return false
	}

	recordTime := accessLogRecord.Time

	if recordTime.After(f.timeStart) && recordTime.Before(f.timeEnd) {
		return false
	}

	return true
}
