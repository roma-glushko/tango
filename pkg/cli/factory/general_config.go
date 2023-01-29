package factory

import (
	"tango/pkg/services/config"

	"github.com/urfave/cli"
)

//
func NewGeneralConfig(cliContext *cli.Context) config.GeneralConfig {
	configFile := cliContext.GlobalString("config-file")
	baseURL := cliContext.GlobalString("base-url")

	return config.NewGeneralConfig(
		configFile,
		baseURL,
	)
}
