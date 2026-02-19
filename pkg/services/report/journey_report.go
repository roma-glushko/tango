package report

import (
	"strings"
	"tango/pkg/entity"
	"tango/pkg/services/config"

	"github.com/google/uuid"
)

// JourneyReportWriter knows how to save journey report
type JourneyReportWriter interface {
	Save(reportPath string, journeyReport map[string]*entity.Journey)
}

// DefaultExcludedURIs are URI patterns excluded from journey places by default
var DefaultExcludedURIs = []string{"/customer/section/load"}

// JourneyReportService knows how to prepare journey reports
type JourneyReportService struct {
	baseURL             string
	excludedURIs        []string
	journeyReportWriter JourneyReportWriter
}

// NewJourneyReportService creates a new instance of the services
func NewJourneyReportService(generalConfig config.GeneralConfig, journeyReportWriter JourneyReportWriter) *JourneyReportService {
	return &JourneyReportService{
		baseURL:             generalConfig.BaseURL,
		excludedURIs:        DefaultExcludedURIs,
		journeyReportWriter: journeyReportWriter,
	}
}

// getUUID retrieves unique ID for journey places
func getUUID() string {
	return uuid.New().String()
}

// GenerateReport processes access logs and determine visitor's journeys on the website
func (u *JourneyReportService) GenerateReport(reportPath string, accessRecords <-chan []entity.AccessLogRecord) {
	journeyReport := make(map[string]*entity.Journey, 0)

	for batch := range accessRecords {
		for _, accessRecord := range batch {
			for _, ip := range accessRecord.SplitIPs() {
				if _, ok := journeyReport[ip]; !ok {
					journeyReport[ip] = &entity.Journey{
						ID: getUUID(),
						IP: ip,
					}
				}

				u.addPlace(journeyReport[ip], accessRecord)
			}
		}
	}

	u.journeyReportWriter.Save(reportPath, journeyReport)
}

// addPlace
func (u *JourneyReportService) addPlace(journey *entity.Journey, accessLogRecord entity.AccessLogRecord) {
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
				ResponseCode:  200,   // assume that previous request was successful
				ResponseSize:  0,     // hard to say about size, keep 0 bytes
				RequestMethod: "GET", // usually GET method is cachable, so assume it was used for this request as well
				RefererURL:    "-",
			},
		})

		if lastAddedPlace != nil {
			journey.AddRoad(lastAddedPlace, refererPlace)
		}
	}

	if !u.isExcludedURI(accessLogRecord.URI) {
		currentPlace := journey.AddPlace(&entity.JourneyPlace{
			ID:        getUUID(),
			WasLogged: true,
			Data:      &accessLogRecord,
		})

		journey.AddRoad(refererPlace, currentPlace)
	}
}

// isExcludedURI checks if a URI should be excluded from journey places
func (u *JourneyReportService) isExcludedURI(uri string) bool {
	for _, excluded := range u.excludedURIs {
		if strings.Contains(uri, excluded) {
			return true
		}
	}
	return false
}

// isInternalReferer checks if given URL is internal referer link
func (u *JourneyReportService) isInternalReferer(refererURL string) bool {
	return strings.HasPrefix(refererURL, u.baseURL)
}
