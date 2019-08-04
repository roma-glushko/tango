package di

import (
	"tango/internal/cli/component"
	"tango/internal/cli/factory"
	"tango/internal/infrastructure/reader"
	"tango/internal/infrastructure/writer"
	"tango/internal/usecase"
	"tango/internal/usecase/config"
	"tango/internal/usecase/filter"
	"tango/internal/usecase/report"

	"github.com/urfave/cli"
)

//
func InitGeneralConfig(cliContext *cli.Context) config.GeneralConfig {
	return factory.NewGeneralConfig(cliContext)
}

//
func InitFilterConfig(cliContext *cli.Context) config.FilterConfig {
	return factory.NewFilterConfig(cliContext)
}

//
func InitFilterAccessLogUsecase(filterConfig config.FilterConfig) usecase.FilterAccessLogUsecase {
	filters := []usecase.AccessLogFilter{
		filter.NewIPFilter(filterConfig),
		filter.NewTimeFilter(filterConfig),
		filter.NewUrlFilter(filterConfig),
		filter.NewAssetFilter(filterConfig),
		filter.NewUserAgentFilter(filterConfig),
	}

	return usecase.NewFilterAccessLogUsecase(filters)
}

//
func InitReadAccessLogUsecase(filterConfig config.FilterConfig) *usecase.ReadAccessLogUsecase {
	accessLogReader := reader.NewAccessLogReader()
	readerProgressDecorator := component.NewReaderProgressDecorator(accessLogReader)
	filterAccessLogUsecase := InitFilterAccessLogUsecase(filterConfig)

	return usecase.NewReadAccessLogUsecase(readerProgressDecorator, filterAccessLogUsecase)
}

//
func InitGeoReportUsecase() *report.GeoReportUsecase {
	geoReportWriter := writer.NewGeoReportCsvWriter()

	return report.NewGeoReportUsecase(geoReportWriter)
}

//
func InitCustomReportUsecase() *report.CustomReportUsecase {
	customReportWriter := writer.NewCustomReportCsvWriter()

	return report.NewCustomReportUsecase(customReportWriter)
}
