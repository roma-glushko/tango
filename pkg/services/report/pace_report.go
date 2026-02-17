package report

import (
	"tango/pkg/entity"
)

var minuteTimeFormat = "2006-01-02 15:04" // minute time group template
var hourTimeFormat = "2006-01-02 15 h"    // hour time group template

// PaceIpReportItem
type PaceIpReportItem struct {
	IP       string
	Requests uint64
	Browser  string
}

// PaceMinuteReportItem
type PaceMinuteReportItem struct {
	Time     string
	IpPaces  map[string]*PaceIpReportItem
	Requests uint64
}

// PaceHourReportItem
type PaceHourReportItem struct {
	Time            string
	MinutePaceItems []*PaceMinuteReportItem
	Requests        uint64
}

// PaceReportWriter
type PaceReportWriter interface {
	Save(reportPath string, paceReport []*PaceHourReportItem)
}

// PaceReportService
type PaceReportService struct {
	paceReportWriter PaceReportWriter
}

func NewPaceReportService(paceReportWriter PaceReportWriter) *PaceReportService {
	return &PaceReportService{
		paceReportWriter: paceReportWriter,
	}
}

// GenerateReport processes access logs and collects request pace reports
func (u *PaceReportService) GenerateReport(reportPath string, accessRecords []entity.AccessLogRecord) {
	var paceReport = make([]*PaceHourReportItem, 0)

	for _, accessRecord := range accessRecords {
		ipList := accessRecord.IP
		browser := accessRecord.UserAgent
		hourTimeGroup := accessRecord.Time.Format(hourTimeFormat)
		minuteTimeGroup := accessRecord.Time.Format(minuteTimeFormat)

		lastHourReportItem := u.findPaceHourReportItem(&paceReport, hourTimeGroup)
		lastMinuteReportItem := u.findPaceMinuteReportItem(&lastHourReportItem.MinutePaceItems, minuteTimeGroup)

		for _, ip := range ipList {
			if _, found := lastMinuteReportItem.IpPaces[ip]; !found {
				lastMinuteReportItem.IpPaces[ip] = &PaceIpReportItem{
					IP:       ip,
					Requests: 0,
					Browser:  browser,
				}
			}

			lastMinuteReportItem.IpPaces[ip].Requests++
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
		IpPaces:  make(map[string]*PaceIpReportItem),
		Requests: 0,
	}
	*items = append(*items, item)

	return item
}
