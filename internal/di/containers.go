package di

import (
	"tango/internal/cli/component"
	"tango/internal/cli/factory"
	"tango/internal/infrastructure/reader"
	"tango/internal/infrastructure/writer"
	"tango/internal/usecase"
	"tango/internal/usecase/config"
	"tango/internal/usecase/filter"
	"tango/internal/usecase/processor"
	"tango/internal/usecase/report"

	"github.com/urfave/cli"
)

//
func InitGeneralConfig(cliContext *cli.Context) config.GeneralConfig {
	return factory.NewGeneralConfig(cliContext)
}

//
func InitProcessorConfig(cliContext *cli.Context) config.ProcessorConfig {
	return factory.NewProcessorConfig(cliContext)
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
func InitIpProcessor(processorConfig config.ProcessorConfig) processor.IPProcessor {
	return processor.NewIPProcessor(processorConfig)
}

//
func InitReadAccessLogUsecase(processorConfig config.ProcessorConfig, filterConfig config.FilterConfig) *usecase.ReadAccessLogUsecase {
	accessLogReader := reader.NewAccessLogReader()
	readerProgressDecorator := component.NewReaderProgressDecorator(accessLogReader)
	ipProcessor := InitIpProcessor(processorConfig)
	filterAccessLogUsecase := InitFilterAccessLogUsecase(filterConfig)

	return usecase.NewReadAccessLogUsecase(readerProgressDecorator, filterAccessLogUsecase, ipProcessor)
}

//
func InitGeoReportUsecase() *report.GeoReportUsecase {
	geoReportWriter := writer.NewGeoReportCsvWriter()

	return report.NewGeoReportUsecase(geoReportWriter)
}

//
func InitBrowserReportUsecase() *report.BrowserReportUsecase {
	browserReportWriter := writer.NewBrowserReportCsvWriter()

	return report.NewBrowserReportUsecase(browserReportWriter)
}

//
func InitRequestReportUsecase() *report.RequestReportUsecase {
	requestReportWriter := writer.NewRequestReportCsvWriter()

	return report.NewRequestReportUsecase(requestReportWriter)
}

// InitJourneyReportUsecase inits a usecase
func InitJourneyReportUsecase(generalConfig config.GeneralConfig) *report.JourneyReportUsecase {
	journeyReportWriter := writer.NewJourneyReportHTMLWriter()

	return report.NewJourneyReportUsecase(generalConfig, journeyReportWriter)
}

// InitCustomReportUsecase inits a usecase
func InitCustomReportUsecase() *report.CustomReportUsecase {
	customReportWriter := writer.NewCustomReportCsvWriter()

	return report.NewCustomReportUsecase(customReportWriter)
}
