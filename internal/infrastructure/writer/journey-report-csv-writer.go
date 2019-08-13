package writer

import (
	"encoding/csv"
	"log"
	"os"
	"tango/internal/usecase/report"
)

type JourneyReportCsvWriter struct {
}

//
func NewJourneyReportCsvWriter() *JourneyReportCsvWriter {
	return &JourneyReportCsvWriter{}
}

// Save journey report to CSV file
func (w *JourneyReportCsvWriter) Save(filePath string, journeyReport map[string]*report.JourneyReport) {
	file, err := os.Create(filePath)

	if err != nil {
		log.Fatal("Error on writing visitor's journey report: ", err)
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Header
	writer.Write([]string{})

	// Body
	for _, requestReportItem := range journeyReport {
		err := writer.Write([]string{})

		if err != nil {
			log.Fatal("Error on writing visitor's journey report: ", err)
		}
	}
}
