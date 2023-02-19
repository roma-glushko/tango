package report

import (
	"strings"
	"sync"
	entity "tango/pkg/entity"
	"tango/pkg/services/mapper"
)

//
type BrowserReportItem struct {
	Browser    string
	Category   string
	SampleUrl  string
	Requests   uint64
	Bandwidth  uint64
	UserAgents map[string]bool
}

//
type BrowserReportWriter interface {
	Save(reportPath string, browserReport map[string]*BrowserReportItem)
}

//
type BrowserReportService struct {
	logMapper           *mapper.AccessLogMapper
	browserReportWriter BrowserReportWriter
}

//
func NewBrowserReportService(logMapper *mapper.AccessLogMapper, browserReportWriter BrowserReportWriter) *BrowserReportService {
	return &BrowserReportService{
		logMapper:           logMapper,
		browserReportWriter: browserReportWriter,
	}
}

// Process access logs and collect browser reports
func (s *BrowserReportService) GenerateReport(reportPath string, logChan <-chan entity.AccessLogRecord) {
	var browserCategories = entity.GetBrowserClassification()

	var browserReport = make(map[string]*BrowserReportItem)
	var mutex sync.Mutex

	var waitGroup sync.WaitGroup

	for i := 0; i < 4; i++ {
		waitGroup.Add(1)

		go func() {
			defer waitGroup.Done()

			for accessRecord := range logChan {
				userAgent := accessRecord.UserAgent

				browserAgent := userAgent
				browserCategory := "Unknown"

				// classify the current browser
				for _, browserCategoryItem := range browserCategories {
					if strings.Contains(userAgent, browserCategoryItem.Agent) {
						browserAgent = browserCategoryItem.Agent
						browserCategory = browserCategoryItem.Category

						break
					}
				}

				mutex.Lock()

				if _, ok := browserReport[browserAgent]; ok {
					browserReport[browserAgent].Requests++
					browserReport[browserAgent].Bandwidth += accessRecord.ResponseSize

					// add a new unique occurance of user agent
					if _, found := browserReport[browserAgent].UserAgents[userAgent]; !found {
						browserReport[browserAgent].UserAgents[userAgent] = true
					}

					s.logMapper.Recycle(accessRecord)
					mutex.Unlock()
					continue
				}

				browserReport[browserAgent] = &BrowserReportItem{
					Browser:    browserAgent,
					Category:   browserCategory,
					SampleUrl:  accessRecord.URI,
					Requests:   1,
					Bandwidth:  accessRecord.ResponseSize,
					UserAgents: map[string]bool{userAgent: true},
				}
				s.logMapper.Recycle(accessRecord)
				mutex.Unlock()
			}
		}()
	}

	waitGroup.Wait()

	s.browserReportWriter.Save(reportPath, browserReport)
}
