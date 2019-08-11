package report

import (
	"strings"
	"tango/internal/domain/entity"
)

//
type BrowserReportItem struct {
	Browser    string
	Category   string
	SampleUrl  string
	Requests   uint64
	Bandwith   uint64
	UserAgents map[string]bool
}

//
type BrowserReportWriter interface {
	Save(reportPath string, browserReport map[string]*BrowserReportItem)
}

//
type BrowserReportUsecase struct {
	browserReportWriter BrowserReportWriter
}

//
func NewBrowserReportUsecase(browserReportWriter BrowserReportWriter) *BrowserReportUsecase {
	return &BrowserReportUsecase{
		browserReportWriter: browserReportWriter,
	}
}

// Process access logs and collect browser reports
func (u *BrowserReportUsecase) GenerateReport(reportPath string, accessRecords []entity.AccessLogRecord) {
	var browserReport = make(map[string]*BrowserReportItem)
	var browserCategories = entity.GetBrowserClassification()

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
			browserReport[browserAgent].Bandwith += accessRecord.ResponseSize

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
			Bandwith:   accessRecord.ResponseSize,
			UserAgents: map[string]bool{userAgent: true},
		}
	}

	u.browserReportWriter.Save(reportPath, browserReport)
}
