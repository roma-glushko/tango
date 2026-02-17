package writer

import (
	"log"
	"os"
	"strconv"
	"tango/pkg/services/report"
)

// RequestReportHeaders defines CSV header columns for the request report.
var RequestReportHeaders = []string{
	"Path",
	"Requests",
	"Response Code",
	"Referer URLs",
}

// RequestReportCsvWriter writes request reports to CSV files.
type RequestReportCsvWriter struct {
}

// NewRequestReportCsvWriter creates a new RequestReportCsvWriter instance.
func NewRequestReportCsvWriter() *RequestReportCsvWriter {
	return &RequestReportCsvWriter{}
}

// Save writes a request report to a CSV file.
func (w *RequestReportCsvWriter) Save(filePath string, requestReport map[string]*report.RequestReportItem) {
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal("Error on writing request report: ", err)
	}

	defer func() { _ = file.Close() }()

	writer, buffered := newBufferedCsvWriter(file)
	defer buffered.Flush()
	defer writer.Flush()

	// Header
	if err := writer.Write(RequestReportHeaders); err != nil {
		log.Println(err)
		return
	}

	// Body
	for _, requestReportItem := range requestReport {
		err := writer.Write([]string{
			requestReportItem.Path,
			strconv.FormatUint(requestReportItem.Requests, 10),
			strconv.FormatUint(requestReportItem.ResponseCode, 10),
			newLineSeparated(requestReportItem.RefererURLs),
		})

		if err != nil {
			log.Println("Error on writing request report: ", err)
			return
		}
	}
}
