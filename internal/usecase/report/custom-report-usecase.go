package report

import (
	"tango/internal/domain/entity"
)

//
type CustomReportWriter interface {
	Save(reportPath string, accessLogs []entity.AccessLogRecord)
}

//
type CustomReportUsecase struct {
	customReportWriter CustomReportWriter
}

//
func NewCustomReportUsecase(customReportWriter CustomReportWriter) *CustomReportUsecase {
	return &CustomReportUsecase{
		customReportWriter: customReportWriter,
	}
}

// Save a custom log based on access logs
func (u *CustomReportUsecase) GenerateReport(reportPath string, accessRecords []entity.AccessLogRecord) {
	// nothing to do here yet
	u.customReportWriter.Save(reportPath, accessRecords)
}
