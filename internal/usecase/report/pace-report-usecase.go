package report

import (
	"tango/internal/domain/entity"
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

// PaceReportUsecase
type PaceReportUsecase struct {
	paceReportWriter PaceReportWriter
}

//
func NewPaceReportUsecase(paceReportWriter PaceReportWriter) *PaceReportUsecase {
	return &PaceReportUsecase{
		paceReportWriter: paceReportWriter,
	}
}

// GenerateReport processes access logs and collects request pace reports
func (u *PaceReportUsecase) GenerateReport(reportPath string, accessRecords []entity.AccessLogRecord) {
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

func (u *PaceReportUsecase) findPaceHourReportItem(paceHourReport *[]*PaceHourReportItem, time string) *PaceHourReportItem {
	lastPaceItemIndex := len(*paceHourReport) - 1

	if lastPaceItemIndex < 0 {
		lastReportItem := &PaceHourReportItem{
			Time:            time,
			MinutePaceItems: make([]*PaceMinuteReportItem, 0),
			Requests:        0,
		}

		*paceHourReport = append(*paceHourReport, lastReportItem)

		return lastReportItem
	}

	lastReportItem := (*paceHourReport)[lastPaceItemIndex]

	if lastReportItem.Time != time {
		lastReportItem := &PaceHourReportItem{
			Time:            time,
			MinutePaceItems: make([]*PaceMinuteReportItem, 0),
			Requests:        0,
		}

		*paceHourReport = append(*paceHourReport, lastReportItem)

		return lastReportItem
	}

	return lastReportItem
}

func (u *PaceReportUsecase) findPaceMinuteReportItem(paceReport *[]*PaceMinuteReportItem, time string) *PaceMinuteReportItem {
	lastPaceItemIndex := len(*paceReport) - 1

	if lastPaceItemIndex < 0 {
		lastReportItem := &PaceMinuteReportItem{
			Time:     time,
			IpPaces:  make(map[string]*PaceIpReportItem),
			Requests: 0,
		}

		*paceReport = append(*paceReport, lastReportItem)

		return lastReportItem
	}

	lastReportItem := (*paceReport)[lastPaceItemIndex]

	if lastReportItem.Time != time {
		lastReportItem := &PaceMinuteReportItem{
			Time:     time,
			IpPaces:  make(map[string]*PaceIpReportItem),
			Requests: 0,
		}

		*paceReport = append(*paceReport, lastReportItem)

		return lastReportItem
	}

	return lastReportItem
}
