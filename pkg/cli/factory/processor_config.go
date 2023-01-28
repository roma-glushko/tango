package factory

import (
	"tango/pkg/services/config"

	"github.com/urfave/cli"
)

//
func NewProcessorConfig(cliContext *cli.Context) config.ProcessorConfig {
	systemIpList := cliContext.GlobalStringSlice("system-ips")

	return config.NewProcessorConfig(
		systemIpList,
	)
}
