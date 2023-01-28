package factory

import (
	"tango/pkg/services/config"

	"github.com/urfave/cli"
)

//
func NewReportConfig(cliContext *cli.Context) config.ReportConfig {
	logFile := cliContext.String("log-file")
	reportFile := cliContext.String("report-file")

	return config.NewReportConfig(
		logFile,
		reportFile,
	)
}
