package di

import (
	"tango/pkg/adapters/filesystem"
	"tango/pkg/adapters/geo"
	"tango/pkg/adapters/reader"
	"tango/pkg/adapters/writer"
	"tango/pkg/cli/factory"
	"tango/pkg/services"
	"tango/pkg/services/config"
	"tango/pkg/services/filter"
	"tango/pkg/services/geodata"
	"tango/pkg/services/mapper"
	"tango/pkg/services/processor"
	"tango/pkg/services/report"

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

// InitFilterAccessLogService
func InitFilterAccessLogService(filterConfig config.FilterConfig) services.FilterAccessLogService {
	filters := []services.AccessLogFilter{
		filter.NewIPFilter(filterConfig),
		filter.NewTimeFilter(filterConfig),
		filter.NewUrlFilter(filterConfig),
		filter.NewAssetFilter(filterConfig),
		filter.NewUserAgentFilter(filterConfig),
	}

	return services.NewFilterAccessLogService(filters)
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

// InitInstallMaxmindLibService inits a services
func InitInstallMaxmindLibService() *geodata.InstallMaxmindLibService {
	return geodata.NewInstallMaxmindLibService()
}

// InitGenerateMaxmindConfService inits a services
func InitGenerateMaxmindConfService() *geodata.GenerateMaxmindConfService {
	return geodata.NewGenerateMaxmindConfService()
}

// InitReadAccessLogService inits a services
func InitReadAccessLogService(logMapper *mapper.AccessLogMapper, processorConfig config.ProcessorConfig, filterConfig config.FilterConfig) *services.ReadAccessLogService {
	logReader := reader.NewAccessLogReader()
	readProgress := reader.NewReadProgress()

	ipProcessor := InitIpProcessor(processorConfig)
	filterAccessLogService := InitFilterAccessLogService(filterConfig)

	return services.NewReadAccessLogService(
		logMapper,
		logReader,
		readProgress,
		filterAccessLogService,
		ipProcessor,
	)
}

// InitGeoReportService inits a services
func InitGeoReportService(logMapper *mapper.AccessLogMapper, maxmindGeoLibPath string) *report.GeoReportService {
	geoProvider := geo.NewMaxMindGeoProvider(maxmindGeoLibPath)
	geoReportWriter := writer.NewGeoReportCsvWriter()

	return report.NewGeoReportService(logMapper, geoProvider, geoReportWriter)
}

// InitBrowserReportService inits a services
func InitBrowserReportService(logMapper *mapper.AccessLogMapper) *report.BrowserReportService {
	browserReportWriter := writer.NewBrowserReportCsvWriter()

	return report.NewBrowserReportService(logMapper, browserReportWriter)
}

// InitRequestReportService inits a services
func InitRequestReportService(logMapper *mapper.AccessLogMapper) *report.RequestReportService {
	requestReportWriter := writer.NewRequestReportCsvWriter()

	return report.NewRequestReportService(logMapper, requestReportWriter)
}

// InitPaceReportService inits a services
func InitPaceReportService(logMapper *mapper.AccessLogMapper) *report.PaceReportService {
	paceReportWriter := writer.NewPaceReportCsvWriter()

	return report.NewPaceReportService(logMapper, paceReportWriter)
}

// InitJourneyReportService inits a services
func InitJourneyReportService(logMapper *mapper.AccessLogMapper, generalConfig config.GeneralConfig) *report.JourneyReportService {
	journeyReportWriter := writer.NewJourneyReportHTMLWriter()

	return report.NewJourneyReportService(logMapper, generalConfig, journeyReportWriter)
}

// InitCustomReportService inits a services
func InitCustomReportService(logMapper *mapper.AccessLogMapper) *report.CustomReportService {
	customReportWriter := writer.NewCustomReportCsvWriter(logMapper)

	return report.NewCustomReportService(customReportWriter)
}
