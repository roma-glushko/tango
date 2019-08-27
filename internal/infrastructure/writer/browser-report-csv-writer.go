package writer

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"tango/internal/usecase/report"
)

//
type BrowserReportCsvWriter struct {
}

//
func NewBrowserReportCsvWriter() *BrowserReportCsvWriter {
	return &BrowserReportCsvWriter{}
}

//
func byteCountDecimal(b uint64) string {
	const unit = 1000

	if b < unit {
		return fmt.Sprintf("%d B", b)
	}

	div, exp := int64(unit), 0

	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}

	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "kMGTPE"[exp])
}

// todo: remove duplicated code
func newLineSeparated(boolMap map[string]bool) string {
	result := ""

	if len(boolMap) == 0 {
		return result
	}

	for userAgent := range boolMap {
		result += userAgent + "\n"
	}

	return result
}

// Save Browser Report to CSV file
func (w *BrowserReportCsvWriter) Save(reportPath string, browserReport map[string]*report.BrowserReportItem) {
	file, err := os.Create(reportPath)

	if err != nil {
		log.Fatal("Error on writing browser report: ", err)
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Header
	writer.Write([]string{
		"Category",
		"Browser",
		"Requests",
		"Bandwith",
		"Sample URL",
		"User Agents",
	})

	// Body
	for _, browserReportItem := range browserReport {
		err := writer.Write([]string{
			browserReportItem.Category,
			browserReportItem.Browser,
			strconv.FormatUint(browserReportItem.Requests, 10),
			byteCountDecimal(browserReportItem.Bandwith),
			browserReportItem.SampleUrl,
			newLineSeparated(browserReportItem.UserAgents),
		})

		if err != nil {
			log.Fatal("Error on writing geolocation report: ", err)
		}
	}
}
