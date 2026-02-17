package report

import (
	"tango/pkg/entity"
)

// CustomReportWriter is an interface for saving custom reports.
type CustomReportWriter interface {
	Save(reportPath string, accessLogs <-chan []entity.AccessLogRecord)
}

// CustomReportService knows how to generate custom reports.
type CustomReportService struct {
	customReportWriter CustomReportWriter
}

// NewCustomReportService creates a new CustomReportService instance.
func NewCustomReportService(customReportWriter CustomReportWriter) *CustomReportService {
	return &CustomReportService{
		customReportWriter: customReportWriter,
	}
}

// GenerateReport saves a custom log based on access logs.
func (u *CustomReportService) GenerateReport(reportPath string, accessRecords <-chan []entity.AccessLogRecord) {
	u.customReportWriter.Save(reportPath, accessRecords)
}
