package report

import (
	"tango/pkg/entity"
)

//
type CustomReportWriter interface {
	Save(reportPath string, logChan <-chan entity.AccessLogRecord)
}

//
type CustomReportService struct {
	customReportWriter CustomReportWriter
}

//
func NewCustomReportService(customReportWriter CustomReportWriter) *CustomReportService {
	return &CustomReportService{
		customReportWriter: customReportWriter,
	}
}

// Save a custom log based on access logs
func (u *CustomReportService) GenerateReport(reportPath string, logChan <-chan entity.AccessLogRecord) {
	// nothing to do here yet
	u.customReportWriter.Save(reportPath, logChan)
}
