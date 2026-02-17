package factory

import (
	"tango/pkg/services/config"

	"github.com/urfave/cli"
)

// NewProcessorConfig creates a ProcessorConfig from CLI context flags.
func NewProcessorConfig(cliContext *cli.Context) config.ProcessorConfig {
	systemIPList := cliContext.GlobalStringSlice("system-ips")

	return config.NewProcessorConfig(
		systemIPList,
	)
}
