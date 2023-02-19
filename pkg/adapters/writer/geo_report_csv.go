package writer

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"tango/pkg/services/report"
)

var GeoReportHeader = []string{
	"IP",
	"Country",
	"City",
	"Continent",
	"Sample Request",
	"Browser Agent",
	"Count of Requests",
}

type GeoReportCsvWriter struct {
}

//
func NewGeoReportCsvWriter() *GeoReportCsvWriter {
	return &GeoReportCsvWriter{}
}

// Save GeoLocation Report to CSV file
func (w *GeoReportCsvWriter) Save(filePath string, geoReport *report.GeoReport) {
	file, err := os.Create(filePath)

	if err != nil {
		log.Fatal("Error on writing geo report: ", err)
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Header
	if err := writer.Write(GeoReportHeader); err != nil {
		log.Fatal(err)
	}

	// Body
	for ip, geoRecord := range geoReport.Report() {
		err := writer.Write([]string{
			ip,
			geoRecord.GeoData.Country,
			geoRecord.GeoData.City,
			geoRecord.GeoData.Continent,
			geoRecord.SampleRequest,
			geoRecord.BrowserAgent,
			strconv.FormatUint(geoRecord.Requests, 10),
		})

		if err != nil {
			log.Fatal("Error on writing geo report: ", err)
		}
	}
}
