package report

import (
	"strings"
	"sync"
	entity2 "tango/pkg/entity"
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
	browserReportWriter BrowserReportWriter
}

//
func NewBrowserReportService(browserReportWriter BrowserReportWriter) *BrowserReportService {
	return &BrowserReportService{
		browserReportWriter: browserReportWriter,
	}
}

// Process access logs and collect browser reports
func (u *BrowserReportService) GenerateReport(reportPath string, logChan <-chan entity2.AccessLogRecord) {
	var browserCategories = entity2.GetBrowserClassification()

	var browserReport = make(map[string]*BrowserReportItem)
	var mutex sync.Mutex // TODO: try to use sync.Map

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
				mutex.Unlock()
			}
		}()
	}

	waitGroup.Wait()

	u.browserReportWriter.Save(reportPath, browserReport)
}
