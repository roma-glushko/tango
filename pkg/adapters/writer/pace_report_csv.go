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
}

// NewPaceReportCsvWriter creates a new PaceReportCsvWriter instance.
func NewPaceReportCsvWriter() *PaceReportCsvWriter {
	return &PaceReportCsvWriter{}
}

// Save writes a pace report to a CSV file.
func (w *PaceReportCsvWriter) Save(filePath string, paceReport []*report.PaceHourReportItem) {
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal("Error on writing request report: ", err)
	}

	defer func() { _ = file.Close() }()

	writer, buffered := newBufferedCsvWriter(file)
	defer buffered.Flush()
	defer writer.Flush()

	// Header
	if err := writer.Write(PaceReportHeader); err != nil {
		log.Println(err)
		return
	}

	// Body
	for _, hourPaceItem := range paceReport {
		// render minute interval header
		err := writer.Write([]string{
			hourPaceItem.Time,
			"",
			"",
			"",
			"",
			strconv.FormatUint(hourPaceItem.Requests, 10),
		})

		if err != nil {
			log.Println("Error on writing request report: ", err)
			return
		}

		for _, paceMinuteItem := range hourPaceItem.MinutePaceItems {
			// render minute interval header
			err := writer.Write([]string{
				"",
				paceMinuteItem.Time,
				"",
				"",
				strconv.FormatUint(paceMinuteItem.Requests, 10),
				"",
			})

			if err != nil {
				log.Println("Error on writing request report: ", err)
				return
			}

			for ip, ipPaceItem := range paceMinuteItem.IPPaces {
				// render ip paces
				err = writer.Write([]string{
					"",
					"",
					ip,
					ipPaceItem.Browser,
					strconv.FormatUint(ipPaceItem.Requests, 10),
					"",
				})

				if err != nil {
					log.Println("Error on writing request report: ", err)
					return
				}
			}

			// render minute interval summary footer
			err = writer.Write([]string{
				"",
				paceMinuteItem.Time,
				"",
				"",
				strconv.FormatUint(paceMinuteItem.Requests, 10),
				"",
			})

			if err != nil {
				log.Println("Error on writing request report: ", err)
				return
			}
		}

		// render hour interval summary footer
		err = writer.Write([]string{
			hourPaceItem.Time,
			"",
			"",
			"",
			"",
			strconv.FormatUint(hourPaceItem.Requests, 10),
		})

		if err != nil {
			log.Println("Error on writing request report: ", err)
			return
		}
	}
}
