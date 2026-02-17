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
	writeBufferSize int
}

// NewGeoReportCsvWriter creates a new GeoReportCsvWriter instance.
func NewGeoReportCsvWriter(writeBufferSize int) *GeoReportCsvWriter {
	return &GeoReportCsvWriter{writeBufferSize: writeBufferSize}
}

// Save writes a geo location report to a CSV file.
func (w *GeoReportCsvWriter) Save(filePath string, geolocationReport map[string]*report.Geolocation) {
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal("Error on writing geo report: ", err)
	}

	defer func() { _ = file.Close() }()

	writer, buffered := newBufferedCsvWriter(file, w.writeBufferSize)
	defer func() {
		if err := buffered.Flush(); err != nil {
			log.Println("Error flushing geo report buffer: ", err)
		}
	}()
	defer writer.Flush()

	// Header
	if err := writer.Write(GeoReportHeader); err != nil {
		log.Println(err)
		return
	}

	// Body
	row := make([]string, 7)
	for ip, geoLocation := range geolocationReport {
		row[0] = ip
		row[1] = geoLocation.GeoData.Country
		row[2] = geoLocation.GeoData.City
		row[3] = geoLocation.GeoData.Continent
		row[4] = geoLocation.SampleRequest
		row[5] = geoLocation.BrowserAgent
		row[6] = strconv.FormatUint(geoLocation.Requests, 10)

		if err := writer.Write(row); err != nil {
			log.Println("Error on writing geo report: ", err)
			return
		}
	}
}
