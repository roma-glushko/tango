package report

import (
	"crypto/rand"
	"fmt"
	"log"
	"strings"
	"sync"
	entity2 "tango/pkg/entity"
	"tango/pkg/services/config"
)

// JourneyReportWriter knows how to save journey report
type JourneyReportWriter interface {
	Save(reportPath string, journeyReport map[string]*entity2.Journey)
}

// JourneyReportService knows how to prepare journey reports
type JourneyReportService struct {
	baseURL             string
	journeyReportWriter JourneyReportWriter
}

// NewJourneyReportService creates a new instance of the services
func NewJourneyReportService(generalConfig config.GeneralConfig, journeyReportWriter JourneyReportWriter) *JourneyReportService {
	return &JourneyReportService{
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
func (u *JourneyReportService) GenerateReport(reportPath string, logChan <-chan entity2.AccessLogRecord) {
	journeyReport := make(map[string]*entity2.Journey, 0)
	var mutex sync.Mutex // TODO: try to use sync.Map
	var waitGroup sync.WaitGroup

	for i := 0; i < 4; i++ {
		waitGroup.Add(1)

		go func() {
			defer waitGroup.Done()

			for accessRecord := range logChan {
				ipList := accessRecord.IP

				for _, ip := range ipList {
					mutex.Lock()
					if _, ok := journeyReport[ip]; !ok {
						journeyReport[ip] = &entity2.Journey{
							ID: getUUID(),
							IP: ip,
						}
					}

					u.addPlace(journeyReport[ip], accessRecord)
					mutex.Unlock()
				}
			}
		}()
	}

	waitGroup.Wait()

	u.journeyReportWriter.Save(reportPath, journeyReport)
}

// addPlace
func (u *JourneyReportService) addPlace(journey *entity2.Journey, accessLogRecord entity2.AccessLogRecord) {
	refererURI := accessLogRecord.RefererURL

	if u.isInternalReferer(refererURI) {
		refererURI = strings.ReplaceAll(refererURI, u.baseURL, "")
	}

	// try to find referer place in journey
	refererPlace := journey.FindLastPlace(refererURI)

	if refererPlace == nil {
		lastAddedPlace := journey.GetLastPlace()

		refererPlace = journey.AddPlace(&entity2.JourneyPlace{
			ID:        getUUID(),
			WasLogged: false,
			Data: &entity2.AccessLogRecord{
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

	if !strings.Contains(accessLogRecord.URI, "/customer/section/load") { // todo: refactor and remove hardcoded uri
		currentPlace := journey.AddPlace(&entity2.JourneyPlace{
			ID:        getUUID(),
			WasLogged: true,
			Data:      &accessLogRecord,
		})

		journey.AddRoad(refererPlace, currentPlace)
	}
}

// isInternalReferer checks if given URL is internal referer link
func (u *JourneyReportService) isInternalReferer(refererURL string) bool {
	return strings.HasPrefix(refererURL, u.baseURL)
}
