package tmpl

import _ "embed"

// JourneyReportTemplate contains the HTML template for journey reports.
//
//go:embed journey-report.tmpl
var JourneyReportTemplate string
