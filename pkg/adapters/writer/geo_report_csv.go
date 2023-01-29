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
func (w *GeoReportCsvWriter) Save(filePath string, geolocationReport map[string]*report.Geolocation) {
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
	for ip, geoLocation := range geolocationReport {
		err := writer.Write([]string{
			ip,
			geoLocation.GeoData.Country,
			geoLocation.GeoData.City,
			geoLocation.GeoData.Continent,
			geoLocation.SampleRequest,
			geoLocation.BrowserAgent,
			strconv.FormatUint(geoLocation.Requests, 10),
		})

		if err != nil {
			log.Fatal("Error on writing geo report: ", err)
		}
	}
}
