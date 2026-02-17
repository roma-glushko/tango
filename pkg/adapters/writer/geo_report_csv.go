package writer

import (
	"log"
	"os"
	"strconv"
	"tango/pkg/services/report"
)

// GeoReportHeader defines CSV header columns for the geo report.
var GeoReportHeader = []string{
	"IP",
	"Country",
	"City",
	"Continent",
	"Sample Request",
	"Browser Agent",
	"Count of Requests",
}

// GeoReportCsvWriter writes geo location reports to CSV files.
type GeoReportCsvWriter struct {
}

// NewGeoReportCsvWriter creates a new GeoReportCsvWriter instance.
func NewGeoReportCsvWriter() *GeoReportCsvWriter {
	return &GeoReportCsvWriter{}
}

// Save writes a geo location report to a CSV file.
func (w *GeoReportCsvWriter) Save(filePath string, geolocationReport map[string]*report.Geolocation) {
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal("Error on writing geo report: ", err)
	}

	defer func() { _ = file.Close() }()

	writer, buffered := newBufferedCsvWriter(file)
	defer buffered.Flush()
	defer writer.Flush()

	// Header
	if err := writer.Write(GeoReportHeader); err != nil {
		log.Println(err)
		return
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
			log.Println("Error on writing geo report: ", err)
			return
		}
	}
}
