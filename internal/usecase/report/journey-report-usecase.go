package report

import (
	"crypto/rand"
	"fmt"
	"log"
	"strings"
	"tango/internal/domain/entity"
	"tango/internal/usecase/config"
)

// JourneyReportWriter knows how to save journey report
type JourneyReportWriter interface {
	Save(reportPath string, journeyReport map[string]*entity.Journey)
}

// JourneyReportUsecase knows how to prepare journey reports
type JourneyReportUsecase struct {
	baseURL             string
	journeyReportWriter JourneyReportWriter
}

// NewJourneyReportUsecase creates a new instance of the usecase
func NewJourneyReportUsecase(generalConfig config.GeneralConfig, journeyReportWriter JourneyReportWriter) *JourneyReportUsecase {
	return &JourneyReportUsecase{
		baseURL:             generalConfig.BaseURL,
		journeyReportWriter: journeyReportWriter,
	}
}

// getUUID retrieves unique ID for journey places
func getUUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)

	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%x%x%x%x%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

// GenerateReport processes access logs and determine visitor's journeys on the website
func (u *JourneyReportUsecase) GenerateReport(reportPath string, accessRecords []entity.AccessLogRecord) {
	journeyReport := make(map[string]*entity.Journey, 0)

	for _, accessRecord := range accessRecords {
		ipList := accessRecord.IP

		for _, ip := range ipList {
			if _, ok := journeyReport[ip]; !ok {
				journeyReport[ip] = &entity.Journey{
					ID: getUUID(),
					IP: ip,
				}
			}

			u.addPlace(journeyReport[ip], accessRecord)
		}
	}

	u.journeyReportWriter.Save(reportPath, journeyReport)
}

// addPlace
func (u *JourneyReportUsecase) addPlace(journey *entity.Journey, accessLogRecord entity.AccessLogRecord) {
	refererURI := accessLogRecord.RefererURL

	if u.isInternalReferer(refererURI) {
		refererURI = strings.ReplaceAll(refererURI, u.baseURL, "")
	}

	// try to find referer place in journey
	refererPlace := journey.FindLastPlace(refererURI)

	if refererPlace == nil {
		lastAddedPlace := journey.GetLastPlace()

		refererPlace = journey.AddPlace(&entity.JourneyPlace{
			ID:        getUUID(),
			WasLogged: false,
			Data: &entity.AccessLogRecord{
				IP:            accessLogRecord.IP,
				URI:           refererURI,
				Time:          accessLogRecord.Time,
				UserAgent:     accessLogRecord.UserAgent,
				Protocol:      accessLogRecord.Protocol,
				ResponseCode:  200,   // assume that previous request was successfull
				ResponseSize:  0,     // hard to say about size, keep 0 bytes
				RequestMethod: "GET", // usually GET method is cachable, so assume it was used for this request as well
				RefererURL:    "-",
			},
		})

		if lastAddedPlace != nil {
			journey.AddRoad(lastAddedPlace, refererPlace)
		}
	}

	if !strings.Contains(accessLogRecord.URI, "/customer/section/load") {
		currentPlace := journey.AddPlace(&entity.JourneyPlace{
			ID:        getUUID(),
			WasLogged: true,
			Data:      &accessLogRecord,
		})

		journey.AddRoad(refererPlace, currentPlace)
	}
}

// isInternalReferer checks if given URL is internal referer link
func (u *JourneyReportUsecase) isInternalReferer(refererURL string) bool {
	return strings.HasPrefix(refererURL, u.baseURL)
}
