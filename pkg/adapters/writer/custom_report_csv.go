package writer

import (
	"log"
	"os"
	"strconv"
	"strings"
	"tango/pkg/entity"
)

var timeFormat = "2006-01-02 15:04:05 -0700" // todo: add localization for US/EU formats

// CustomReportHeader defines CSV header columns for the custom report.
var CustomReportHeader = []string{
	"Time",
	"IP",
	"URI",
	"Referer URL",
	"Response Code",
	"User Agent",
}

// CustomReportCsvWriter writes custom reports to CSV files.
type CustomReportCsvWriter struct {
	writeBufferSize int
}

// NewCustomReportCsvWriter creates a new CustomReportCsvWriter instance.
func NewCustomReportCsvWriter(writeBufferSize int) *CustomReportCsvWriter {
	return &CustomReportCsvWriter{writeBufferSize: writeBufferSize}
}

// Save writes a custom report to a CSV file.
func (w *CustomReportCsvWriter) Save(filePath string, accessLogs []entity.AccessLogRecord) {
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal("Error on writing custom report: ", err)
	}

	defer func() { _ = file.Close() }()

	writer, buffered := newBufferedCsvWriter(file, w.writeBufferSize)
	defer func() {
		if err := buffered.Flush(); err != nil {
			log.Println("Error flushing custom report buffer: ", err)
		}
	}()
	defer writer.Flush()

	// Header
	if err := writer.Write(CustomReportHeader); err != nil {
		log.Println(err)
		return
	}

	// Body
	row := make([]string, 6)
	for _, accessLog := range accessLogs {
		row[0] = accessLog.Time.Format(timeFormat)
		row[1] = strings.Join(accessLog.IP, ", ")
		row[2] = accessLog.URI
		row[3] = accessLog.RefererURL
		row[4] = strconv.FormatUint(accessLog.ResponseCode, 10)
		row[5] = accessLog.UserAgent

		if err := writer.Write(row); err != nil {
			log.Println("Error on writing custom report: ", err)
			return
		}
	}
}
