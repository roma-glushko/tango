package writer

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"tango/internal/usecase/report"
)

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
		log.Fatal("Error on writing geolocation report: ", err)
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Header
	writer.Write([]string{
		"IP",
		"Country",
		"City",
		"Count of Requests",
	})

	// Body
	for ip, geoLocation := range geolocationReport {
		err := writer.Write([]string{
			ip,
			geoLocation.Country,
			geoLocation.City,
			strconv.FormatUint(geoLocation.Requests, 10),
		})

		if err != nil {
			log.Fatal("Error on writing geolocation report: ", err)
		}
	}
}
