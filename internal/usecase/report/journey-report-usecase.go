package report

import (
	"tango/internal/domain/entity"
)

type JourneyReport struct {
}

type JourneyReportWriter interface {
	Save(reportPath string, journeyReport map[string]*JourneyReport)
}

type JourneyReportUsecase struct {
	journeyReportWriter JourneyReportWriter
}

//
func NewJourneyReportUsecase(journeyReportWriter JourneyReportWriter) *JourneyReportUsecase {
	return &JourneyReportUsecase{
		journeyReportWriter: journeyReportWriter,
	}
}

// Process access logs and determine visitor's journeys on the website
func (u *JourneyReportUsecase) GenerateReport(reportPath string, accessRecords []entity.AccessLogRecord) {
	journeyReport := make(map[string]*JourneyReport, 0)

	// todo: what's a usecase here?

	u.journeyReportWriter.Save(reportPath, journeyReport)
}
