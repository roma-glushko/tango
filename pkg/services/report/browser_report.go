package report

import (
	"strings"
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
func (u *BrowserReportService) GenerateReport(reportPath string, accessRecords []entity2.AccessLogRecord) {
	var browserReport = make(map[string]*BrowserReportItem)
	var browserCategories = entity2.GetBrowserClassification()

	for _, accessRecord := range accessRecords {
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

		if _, ok := browserReport[browserAgent]; ok {
			browserReport[browserAgent].Requests++
			browserReport[browserAgent].Bandwidth += accessRecord.ResponseSize

			// add a new unique occurance of user agent
			if _, found := browserReport[browserAgent].UserAgents[userAgent]; !found {
				browserReport[browserAgent].UserAgents[userAgent] = true
			}

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
	}

	u.browserReportWriter.Save(reportPath, browserReport)
}
