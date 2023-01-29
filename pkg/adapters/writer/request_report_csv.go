package writer

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"tango/pkg/services/report"
)

var RequestReportHeaders = []string{
	"Path",
	"Requests",
	"Response Code",
	"Referer URLs",
}

type RequestReportCsvWriter struct {
}

//
func NewRequestReportCsvWriter() *RequestReportCsvWriter {
	return &RequestReportCsvWriter{}
}

// Save request report to CSV file
func (w *RequestReportCsvWriter) Save(filePath string, requestReport map[string]*report.RequestReportItem) {
	file, err := os.Create(filePath)

	if err != nil {
		log.Fatal("Error on writing request report: ", err)
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Header
	if err := writer.Write(RequestReportHeaders); err != nil {
		log.Fatal(err)
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
			log.Fatal("Error on writing request report: ", err)
		}
	}
}
