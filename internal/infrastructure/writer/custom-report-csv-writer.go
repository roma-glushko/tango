package writer

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
	"tango/internal/domain/entity"
)

var timeFormat = "2006-01-02 15:04:05" // todo: add localization for US/EU formats

var CustomReportHeader = []string{
	"Time",
	"IP",
	"URI",
	"Referer URL",
	"Response Code",
	"User Agent",
}

//
type CustomReportCsvWriter struct {
}

//
func NewCustomReportCsvWriter() *CustomReportCsvWriter {
	return &CustomReportCsvWriter{}
}

// Save GeoLocation Report to CSV file
func (w *CustomReportCsvWriter) Save(filePath string, accessLogs []entity.AccessLogRecord) {
	file, err := os.Create(filePath)

	if err != nil {
		log.Fatal("Error on writing custom report: ", err)
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Header
	writer.Write(CustomReportHeader)

	// Body
	for _, accessLog := range accessLogs {
		err := writer.Write([]string{
			accessLog.Time.Format(timeFormat),
			strings.Join(accessLog.IP, ", "),
			accessLog.URI,
			accessLog.RefererURL,
			strconv.FormatUint(accessLog.ResponseCode, 10),
			accessLog.UserAgent,
		})

		if err != nil {
			log.Fatal("Error on writing custom report: ", err)
		}
	}
}
