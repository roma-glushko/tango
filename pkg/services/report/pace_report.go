package report

import (
	"tango/pkg/entity"
)

var minuteTimeFormat = "2006-01-02 15:04" // minute time group template
var hourTimeFormat = "2006-01-02 15 h"    // hour time group template

// PaceIPReportItem represents an IP entry in the pace report.
type PaceIPReportItem struct {
	IP       string
	Requests uint64
	Browser  string
}

// PaceMinuteReportItem represents a minute interval in the pace report.
type PaceMinuteReportItem struct {
	Time     string
	IPPaces  map[string]*PaceIPReportItem
	Requests uint64
}

// PaceHourReportItem represents an hour interval in the pace report.
type PaceHourReportItem struct {
	Time            string
	MinutePaceItems []*PaceMinuteReportItem
	Requests        uint64
}

// PaceReportWriter is an interface for saving pace reports.
type PaceReportWriter interface {
	Save(reportPath string, paceReport []*PaceHourReportItem)
}

// PaceReportService knows how to generate pace reports.
type PaceReportService struct {
	paceReportWriter PaceReportWriter
}

// NewPaceReportService creates a new PaceReportService instance.
func NewPaceReportService(paceReportWriter PaceReportWriter) *PaceReportService {
	return &PaceReportService{
		paceReportWriter: paceReportWriter,
	}
}

// GenerateReport processes access logs and collects request pace reports.
func (u *PaceReportService) GenerateReport(reportPath string, accessRecords <-chan entity.AccessLogRecord) {
	var paceReport = make([]*PaceHourReportItem, 0)

	for accessRecord := range accessRecords {
		browser := accessRecord.UserAgent
		hourTimeGroup := accessRecord.Time.Format(hourTimeFormat)
		minuteTimeGroup := accessRecord.Time.Format(minuteTimeFormat)

		lastHourReportItem := u.findPaceHourReportItem(&paceReport, hourTimeGroup)
		lastMinuteReportItem := u.findPaceMinuteReportItem(&lastHourReportItem.MinutePaceItems, minuteTimeGroup)

		for _, ip := range accessRecord.SplitIPs() {
			if _, found := lastMinuteReportItem.IPPaces[ip]; !found {
				lastMinuteReportItem.IPPaces[ip] = &PaceIPReportItem{
					IP:       ip,
					Requests: 0,
					Browser:  browser,
				}
			}

			lastMinuteReportItem.IPPaces[ip].Requests++
		}

		lastMinuteReportItem.Requests++
		lastHourReportItem.Requests++
	}

	u.paceReportWriter.Save(reportPath, paceReport)
}

func (u *PaceReportService) findPaceHourReportItem(items *[]*PaceHourReportItem, timeGroup string) *PaceHourReportItem {
	if n := len(*items); n > 0 {
		if last := (*items)[n-1]; last.Time == timeGroup {
			return last
		}
	}

	item := &PaceHourReportItem{
		Time:            timeGroup,
		MinutePaceItems: make([]*PaceMinuteReportItem, 0),
		Requests:        0,
	}
	*items = append(*items, item)

	return item
}

func (u *PaceReportService) findPaceMinuteReportItem(items *[]*PaceMinuteReportItem, timeGroup string) *PaceMinuteReportItem {
	if n := len(*items); n > 0 {
		if last := (*items)[n-1]; last.Time == timeGroup {
			return last
		}
	}

	item := &PaceMinuteReportItem{
		Time:     timeGroup,
		IPPaces:  make(map[string]*PaceIPReportItem),
		Requests: 0,
	}
	*items = append(*items, item)

	return item
}
