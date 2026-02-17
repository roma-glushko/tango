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
}

// NewCustomReportCsvWriter creates a new CustomReportCsvWriter instance.
func NewCustomReportCsvWriter() *CustomReportCsvWriter {
	return &CustomReportCsvWriter{}
}

// Save writes a custom report to a CSV file.
func (w *CustomReportCsvWriter) Save(filePath string, accessLogs []entity.AccessLogRecord) {
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal("Error on writing custom report: ", err)
	}

	defer func() { _ = file.Close() }()

	writer, buffered := newBufferedCsvWriter(file)
	defer buffered.Flush()
	defer writer.Flush()

	// Header
	if err := writer.Write(CustomReportHeader); err != nil {
		log.Println(err)
		return
	}

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
			log.Println("Error on writing custom report: ", err)
			return
		}
	}
}
