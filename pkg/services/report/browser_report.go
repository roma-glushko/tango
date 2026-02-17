package report

import (
	"strings"
	"tango/pkg/entity"
)

// BrowserReportItem represents a single browser entry in the browser report.
type BrowserReportItem struct {
	Browser    string
	Category   string
	SampleURL  string
	Requests   uint64
	Bandwidth  uint64
	UserAgents map[string]bool
}

// BrowserReportWriter is an interface for saving browser reports.
type BrowserReportWriter interface {
	Save(reportPath string, browserReport map[string]*BrowserReportItem)
}

// BrowserReportService knows how to generate browser reports.
type BrowserReportService struct {
	browserReportWriter BrowserReportWriter
}

// NewBrowserReportService creates a new BrowserReportService instance.
func NewBrowserReportService(browserReportWriter BrowserReportWriter) *BrowserReportService {
	return &BrowserReportService{
		browserReportWriter: browserReportWriter,
	}
}

// GenerateReport processes access logs and collects browser reports.
func (u *BrowserReportService) GenerateReport(reportPath string, accessRecords <-chan entity.AccessLogRecord) {
	var browserReport = make(map[string]*BrowserReportItem)
	var browserCategories = entity.GetBrowserClassification()

	for accessRecord := range accessRecords {
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

			// add a new unique occurrence of user agent
			if _, found := browserReport[browserAgent].UserAgents[userAgent]; !found {
				browserReport[browserAgent].UserAgents[userAgent] = true
			}

			continue
		}

		browserReport[browserAgent] = &BrowserReportItem{
			Browser:    browserAgent,
			Category:   browserCategory,
			SampleURL:  accessRecord.URI,
			Requests:   1,
			Bandwidth:  accessRecord.ResponseSize,
			UserAgents: map[string]bool{userAgent: true},
		}
	}

	u.browserReportWriter.Save(reportPath, browserReport)
}
