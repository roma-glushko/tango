package report

import (
	"strings"
	"tango/internal/domain/entity"
)

//
type BrowserReportItem struct {
	Browser   string
	Category  string
	SampleUrl string
	Requests  uint64
	Bandwith  uint64
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
	var browserInfoMap = entity.GetBrowserClassification()

	for _, accessRecord := range accessRecords {
		userAgent := accessRecord.UserAgent

		browserEngine := userAgent
		browserCategory := "Unknown"

		// classify the current browser
		for browser, category := range browserInfoMap {
			if strings.Contains(userAgent, browser) {
				browserEngine = browser
				browserCategory = category

				break
			}
		}

		if _, ok := browserReport[browserEngine]; ok {
			browserReport[browserEngine].Requests++
			browserReport[browserEngine].Bandwith += accessRecord.ResponseSize

			continue
		}

		browserReport[browserEngine] = &BrowserReportItem{
			Browser:   browserEngine,
			Category:  browserCategory,
			SampleUrl: accessRecord.URI,
			Requests:  1,
			Bandwith:  accessRecord.ResponseSize,
		}
	}

	u.browserReportWriter.Save(reportPath, browserReport)
}
