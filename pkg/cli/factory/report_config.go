package factory

import (
	"tango/pkg/services/config"

	"github.com/urfave/cli"
)

// NewReportConfig creates a ReportConfig from CLI context flags.
func NewReportConfig(cliContext *cli.Context) config.ReportConfig {
	logFile := cliContext.String("log-file")
	reportFile := cliContext.String("report-file")

	return config.NewReportConfig(
		logFile,
		reportFile,
	)
}
