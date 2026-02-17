package writer

import (
	"log"
	"os"
	"strconv"
	"tango/pkg/services/report"
)

// PaceReportHeader defines CSV header columns for the pace report.
var PaceReportHeader = []string{
	"Hour Group",
	"Minute Group",
	"IP",
	"Browser",
	"Pace (req/min)",
	"Pace (req/hour)",
}

// PaceReportCsvWriter writes pace reports to CSV files.
type PaceReportCsvWriter struct {
	writeBufferSize int
}

// NewPaceReportCsvWriter creates a new PaceReportCsvWriter instance.
func NewPaceReportCsvWriter(writeBufferSize int) *PaceReportCsvWriter {
	return &PaceReportCsvWriter{writeBufferSize: writeBufferSize}
}

// Save writes a pace report to a CSV file.
func (w *PaceReportCsvWriter) Save(filePath string, paceReport []*report.PaceHourReportItem) {
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal("Error on writing request report: ", err)
	}

	defer func() { _ = file.Close() }()

	writer, buffered := newBufferedCsvWriter(file, w.writeBufferSize)
	defer func() {
		if err := buffered.Flush(); err != nil {
			log.Println("Error flushing pace report buffer: ", err)
		}
	}()
	defer writer.Flush()

	// Header
	if err := writer.Write(PaceReportHeader); err != nil {
		log.Println(err)
		return
	}

	// Body
	row := make([]string, 6)
	for _, hourPaceItem := range paceReport {
		// render hour interval header
		row[0] = hourPaceItem.Time
		row[1] = ""
		row[2] = ""
		row[3] = ""
		row[4] = ""
		row[5] = strconv.FormatUint(hourPaceItem.Requests, 10)

		if err := writer.Write(row); err != nil {
			log.Println("Error on writing request report: ", err)
			return
		}

		for _, paceMinuteItem := range hourPaceItem.MinutePaceItems {
			// render minute interval header
			row[0] = ""
			row[1] = paceMinuteItem.Time
			row[2] = ""
			row[3] = ""
			row[4] = strconv.FormatUint(paceMinuteItem.Requests, 10)
			row[5] = ""

			if err := writer.Write(row); err != nil {
				log.Println("Error on writing request report: ", err)
				return
			}

			for ip, ipPaceItem := range paceMinuteItem.IPPaces {
				// render ip paces
				row[0] = ""
				row[1] = ""
				row[2] = ip
				row[3] = ipPaceItem.Browser
				row[4] = strconv.FormatUint(ipPaceItem.Requests, 10)
				row[5] = ""

				if err := writer.Write(row); err != nil {
					log.Println("Error on writing request report: ", err)
					return
				}
			}

			// render minute interval summary footer
			row[0] = ""
			row[1] = paceMinuteItem.Time
			row[2] = ""
			row[3] = ""
			row[4] = strconv.FormatUint(paceMinuteItem.Requests, 10)
			row[5] = ""

			if err := writer.Write(row); err != nil {
				log.Println("Error on writing request report: ", err)
				return
			}
		}

		// render hour interval summary footer
		row[0] = hourPaceItem.Time
		row[1] = ""
		row[2] = ""
		row[3] = ""
		row[4] = ""
		row[5] = strconv.FormatUint(hourPaceItem.Requests, 10)

		if err := writer.Write(row); err != nil {
			log.Println("Error on writing request report: ", err)
			return
		}
	}
}
