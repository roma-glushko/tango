package report

import (
	"sync"
	"tango/pkg/entity"
	"tango/pkg/services/mapper"
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
	logMapper        *mapper.AccessLogMapper
	paceReportWriter PaceReportWriter
}

//
func NewPaceReportService(logMapper *mapper.AccessLogMapper, paceReportWriter PaceReportWriter) *PaceReportService {
	return &PaceReportService{
		logMapper:        logMapper,
		paceReportWriter: paceReportWriter,
	}
}

// GenerateReport processes access logs and collects request pace reports
func (s *PaceReportService) GenerateReport(reportPath string, logChan <-chan entity.AccessLogRecord) {
	var paceReport = make([]*PaceHourReportItem, 0)
	var mutex sync.Mutex // TODO: try to use sync.Map

	var waitGroup sync.WaitGroup

	for i := 0; i < 4; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()

			for accessRecord := range logChan {
				ipList := accessRecord.IP
				browser := accessRecord.UserAgent
				hourTimeGroup := accessRecord.Time.Format(hourTimeFormat)
				minuteTimeGroup := accessRecord.Time.Format(minuteTimeFormat)

				mutex.Lock()
				lastHourReportItem := s.findPaceHourReportItem(&paceReport, hourTimeGroup)
				lastMinuteReportItem := s.findPaceMinuteReportItem(&lastHourReportItem.MinutePaceItems, minuteTimeGroup)

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

				s.logMapper.Recycle(accessRecord)
				mutex.Unlock()
			}
		}()
	}

	waitGroup.Wait()

	s.paceReportWriter.Save(reportPath, paceReport)
}

func (u *PaceReportService) findPaceHourReportItem(paceHourReport *[]*PaceHourReportItem, time string) *PaceHourReportItem {
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

func (u *PaceReportService) findPaceMinuteReportItem(paceReport *[]*PaceMinuteReportItem, time string) *PaceMinuteReportItem {
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
