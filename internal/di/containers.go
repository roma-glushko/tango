package di

import (
	"tango/internal/cli/component"
	"tango/internal/cli/factory"
	"tango/internal/infrastructure/filesystem"
	"tango/internal/infrastructure/geo"
	"tango/internal/infrastructure/reader"
	"tango/internal/infrastructure/writer"
	"tango/internal/usecase"
	"tango/internal/usecase/config"
	"tango/internal/usecase/filter"
	"tango/internal/usecase/geodata"
	"tango/internal/usecase/processor"
	"tango/internal/usecase/report"

	"github.com/urfave/cli"
)

// InitReportConfig inits a config provider
func InitReportConfig(cliContext *cli.Context) config.ReportConfig {
	return factory.NewReportConfig(cliContext)
}

// InitGeneralConfig inits a config provider
func InitGeneralConfig(cliContext *cli.Context) config.GeneralConfig {
	return factory.NewGeneralConfig(cliContext)
}

// InitProcessorConfig inits a config provider
func InitProcessorConfig(cliContext *cli.Context) config.ProcessorConfig {
	return factory.NewProcessorConfig(cliContext)
}

// InitFilterConfig inits a config provider
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

// InitHomeDirResolver inits home dir resolver
func InitHomeDirResolver() *filesystem.HomeDirResolver {
	return filesystem.NewHomeDirResolver()
}

// InitMaxmindGeoLibResolver inits Maxmind Geo Lib resolver
func InitMaxmindGeoLibResolver() *geo.MaxMindGeoLibResolver {
	homeDirResolver := InitHomeDirResolver()

	return geo.NewMaxMindGeoLibResolver(homeDirResolver)
}

// InitMaxmindConfResolver inits Maxmind Conf file resolver
func InitMaxmindConfResolver() *geo.MaxMindConfResolver {
	homeDirResolver := InitHomeDirResolver()

	return geo.NewMaxMindConfResolver(homeDirResolver)
}

// InitInstallMaxmindLibUsecase inits a usecase
func InitInstallMaxmindLibUsecase() *geodata.InstallMaxmindLibUsecase {
	return geodata.NewInstallMaxmindLibUsecase()
}

// InitGenerateMaxmindConfUsecase inits a usecase
func InitGenerateMaxmindConfUsecase() *geodata.GenerateMaxmindConfUsecase {
	return geodata.NewGenerateMaxmindConfUsecase()
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
func InitGeoReportUsecase(maxmindGeoLibPath string) *report.GeoReportUsecase {
	geoProvider := geo.NewMaxMindGeoProvider(maxmindGeoLibPath)
	geoReportWriter := writer.NewGeoReportCsvWriter()

	return report.NewGeoReportUsecase(geoProvider, geoReportWriter)
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
