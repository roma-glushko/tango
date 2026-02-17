package factory

import (
	"tango/pkg/services/config"

	"github.com/urfave/cli"
)

// NewFilterConfig creates a FilterConfig from CLI context flags.
func NewFilterConfig(cliContext *cli.Context) config.FilterConfig {
	assetFilters := cliContext.GlobalStringSlice("asset-filter")
	keepTimeFrames := cliContext.GlobalStringSlice("keep-time-filter")

	uriFilters := cliContext.GlobalStringSlice("uri-filter")
	keepURIFilters := cliContext.GlobalStringSlice("keep-uri-filter")

	ipFilters := cliContext.GlobalStringSlice("ip-filter")
	keepIPFilters := cliContext.GlobalStringSlice("keep-ip-filter")

	uaFilters := cliContext.GlobalStringSlice("ua-filter")
	keepUaFilters := cliContext.GlobalStringSlice("keep-ua-filter")

	return config.NewFilterConfig(
		assetFilters,
		keepTimeFrames,
		uriFilters,
		keepURIFilters,
		ipFilters,
		keepIPFilters,
		uaFilters,
		keepUaFilters,
	)
}
