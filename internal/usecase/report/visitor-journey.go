package report

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"tango/internal/domain/entity"
)

var europeFormat = "2006-01-02 03:04:05"

//
func GetVisitorJourneyReport(accessRecords []entity.AccessLogRecord) map[string][]entity.AccessLogRecord {
	var visitorJourneyReport = make(map[string][]entity.AccessLogRecord)

	for _, accessRecord := range accessRecords {
		var ipList = accessRecord.IP

		// todo: remove system IPs

		for _, ip := range ipList {
			if visitorJourneyReport[ip] == nil {
				visitorJourneyReport[ip] = make([]entity.AccessLogRecord, 0)
			}

			visitorJourneyReport[ip] = append(visitorJourneyReport[ip], accessRecord)
		}
	}

	return visitorJourneyReport
}

// Save GeoLocation Report to CSV file
func SaveVisitorJourneyReport(visitorJourneyReport map[string][]entity.AccessLogRecord, journeyDirPath string) {
	for ip, visitorJorney := range visitorJourneyReport {

		if len(visitorJorney) == 0 {
			continue
		}

		jorneyFile, err := os.Create(journeyDirPath + "/" + ip + ".csv")

		if err != nil {
			log.Fatal("Error on writing visitor journey report: ", err)
		}

		writer := csv.NewWriter(jorneyFile)
		defer writer.Flush()

		// Header
		writer.Write([]string{
			"Time",
			"IP",
			"User Agent",
			"URI",
			"Response Code",
			"Referer URL",
		})

		// Body
		for _, journeyRecord := range visitorJorney {
			err := writer.Write([]string{
				journeyRecord.Time.Format(europeFormat),
				ip,
				journeyRecord.UserAgent,
				journeyRecord.URI,
				strconv.FormatUint(journeyRecord.ResponseCode, 10),
				journeyRecord.RefererURL,
			})

			if err != nil {
				log.Fatal("Error on writing visitor journey report: ", err)
			}
		}
	}
}
