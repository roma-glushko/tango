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
	writeBufferSize int
}

// NewRequestReportCsvWriter creates a new RequestReportCsvWriter instance.
func NewRequestReportCsvWriter(writeBufferSize int) *RequestReportCsvWriter {
	return &RequestReportCsvWriter{writeBufferSize: writeBufferSize}
}

// Save writes a request report to a CSV file.
func (w *RequestReportCsvWriter) Save(filePath string, requestReport map[string]*report.RequestReportItem) {
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal("Error on writing request report: ", err)
	}

	defer func() { _ = file.Close() }()

	writer, buffered := newBufferedCsvWriter(file, w.writeBufferSize)
	defer func() {
		if err := buffered.Flush(); err != nil {
			log.Println("Error flushing request report buffer: ", err)
		}
	}()
	defer writer.Flush()

	// Header
	if err := writer.Write(RequestReportHeaders); err != nil {
		log.Println(err)
		return
	}

	// Body
	row := make([]string, 4)
	for _, requestReportItem := range requestReport {
		row[0] = requestReportItem.Path
		row[1] = strconv.FormatUint(requestReportItem.Requests, 10)
		row[2] = strconv.FormatUint(requestReportItem.ResponseCode, 10)
		row[3] = newLineSeparated(requestReportItem.RefererURLs)

		if err := writer.Write(row); err != nil {
			log.Println("Error on writing request report: ", err)
			return
		}
	}
}
