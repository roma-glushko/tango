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

// InitFilterAccessLogUsecase
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

// InitIpProcessor
func InitIpProcessor(processorConfig config.ProcessorConfig) processor.IPProcessor {
	return processor.NewIPProcessor(processorConfig)
}

// InitReadAccessLogUsecase inits a usecase
func InitReadAccessLogUsecase(processorConfig config.ProcessorConfig, filterConfig config.FilterConfig) *usecase.ReadAccessLogUsecase {
	accessLogReader := reader.NewAccessLogReader()
	readerProgressDecorator := component.NewReaderProgressDecorator(accessLogReader)
	ipProcessor := InitIpProcessor(processorConfig)
	filterAccessLogUsecase := InitFilterAccessLogUsecase(filterConfig)

	return usecase.NewReadAccessLogUsecase(readerProgressDecorator, filterAccessLogUsecase, ipProcessor)
}

// InitGeoReportUsecase inits a usecase
func InitGeoReportUsecase() *report.GeoReportUsecase {
	geoReportWriter := writer.NewGeoReportCsvWriter()

	return report.NewGeoReportUsecase(geoReportWriter)
}

// InitBrowserReportUsecase inits a usecase
func InitBrowserReportUsecase() *report.BrowserReportUsecase {
	browserReportWriter := writer.NewBrowserReportCsvWriter()

	return report.NewBrowserReportUsecase(browserReportWriter)
}

// InitRequestReportUsecase inits a usecase
func InitRequestReportUsecase() *report.RequestReportUsecase {
	requestReportWriter := writer.NewRequestReportCsvWriter()

	return report.NewRequestReportUsecase(requestReportWriter)
}

// InitPaceReportUsecase inits a usecase
func InitPaceReportUsecase() *report.PaceReportUsecase {
	paceReportWriter := writer.NewPaceReportCsvWriter()

	return report.NewPaceReportUsecase(paceReportWriter)
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
