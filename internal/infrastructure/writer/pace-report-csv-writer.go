package writer

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"tango/internal/usecase/report"
)

// PaceReportCsvWriter
type PaceReportCsvWriter struct {
}

// NewPaceReportCsvWriter
func NewPaceReportCsvWriter() *PaceReportCsvWriter {
	return &PaceReportCsvWriter{}
}

// Save stores pace report to CSV file
func (w *PaceReportCsvWriter) Save(filePath string, paceReport []*report.PaceHourReportItem) {
	file, err := os.Create(filePath)

	if err != nil {
		log.Fatal("Error on writing request report: ", err)
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Header
	writer.Write([]string{
		"Hour Group",
		"Minute Group",
		"IP",
		"Browser",
		"Pace (req/min)",
		"Pace (req/hour)",
	})

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
			log.Fatal("Error on writing request report: ", err)
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
				log.Fatal("Error on writing request report: ", err)
			}

			for ip, ipPaceItem := range paceMinuteItem.IpPaces {
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
					log.Fatal("Error on writing request report: ", err)
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
				log.Fatal("Error on writing request report: ", err)
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
			log.Fatal("Error on writing request report: ", err)
		}
	}
}
