package writer

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"tango/pkg/services/report"
)

// BrowserReportHeader defines CSV header columns for the browser report.
var BrowserReportHeader = []string{
	"Category",
	"Browser",
	"Requests",
	"Bandwidth",
	"Sample URL",
	"User Agents",
}

// BrowserReportCsvWriter writes browser reports to CSV files.
type BrowserReportCsvWriter struct {
}

// NewBrowserReportCsvWriter creates a new BrowserReportCsvWriter instance.
func NewBrowserReportCsvWriter() *BrowserReportCsvWriter {
	return &BrowserReportCsvWriter{}
}

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

func newLineSeparated(boolMap map[string]bool) string {
	if len(boolMap) == 0 {
		return ""
	}

	var b strings.Builder

	for key := range boolMap {
		b.WriteString(key)
		b.WriteByte('\n')
	}

	return b.String()
}

// Save writes a browser report to a CSV file.
func (w *BrowserReportCsvWriter) Save(reportPath string, browserReport map[string]*report.BrowserReportItem) {
	file, err := os.Create(reportPath)
	if err != nil {
		log.Fatal("Error on writing browser report: ", err)
	}

	defer func() { _ = file.Close() }()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Header
	if err := writer.Write(BrowserReportHeader); err != nil {
		log.Println("Error on writing browser report header: ", err)
		return
	}

	// Body
	for _, browserReportItem := range browserReport {
		err := writer.Write([]string{
			browserReportItem.Category,
			browserReportItem.Browser,
			strconv.FormatUint(browserReportItem.Requests, 10),
			byteCountDecimal(browserReportItem.Bandwidth),
			browserReportItem.SampleURL,
			newLineSeparated(browserReportItem.UserAgents),
		})

		if err != nil {
			log.Println("Error on writing browser report: ", err)
			return
		}
	}
}
