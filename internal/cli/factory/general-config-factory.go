package factory

import (
	"tango/internal/usecase/config"

	"github.com/urfave/cli"
)

//
func NewGeneralConfig(cliContext *cli.Context) config.GeneralConfig {
	logFile := cliContext.GlobalString("log-file")
	reportFile := cliContext.GlobalString("report-file")
	configFile := cliContext.GlobalString("config-file")

	return config.NewGeneralConfig(
		logFile,
		reportFile,
		configFile,
	)
}
