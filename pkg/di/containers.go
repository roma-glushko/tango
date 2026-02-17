package di

import (
	"tango/pkg/adapters/filesystem"
	"tango/pkg/adapters/geo"
	"tango/pkg/adapters/reader"
	"tango/pkg/adapters/writer"
	"tango/pkg/cli/component"
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

// InitFilterAccessLogService creates a FilterAccessLogService with all filters.
func InitFilterAccessLogService(filterConfig config.FilterConfig) services.FilterAccessLogService {
	filters := []services.AccessLogFilter{
		filter.NewIPFilter(filterConfig),
		filter.NewTimeFilter(filterConfig),
		filter.NewURLFilter(filterConfig),
		filter.NewAssetFilter(filterConfig),
		filter.NewUserAgentFilter(filterConfig),
	}

	return services.NewFilterAccessLogService(filters)
}

// InitIPProcessor creates an IPProcessor from the given config.
func InitIPProcessor(processorConfig config.ProcessorConfig) (processor.IPProcessor, error) {
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

// InitPipelineConfig inits pipeline config
func InitPipelineConfig(cliContext *cli.Context) config.PipelineConfig {
	return factory.NewPipelineConfig(cliContext)
}

// InitReadAccessLogService inits a services
func InitReadAccessLogService(processorConfig config.ProcessorConfig, filterConfig config.FilterConfig, pipelineConfig config.PipelineConfig) (*services.ReadAccessLogService, error) {
	accessLogReader := reader.NewAccessLogReader(pipelineConfig.ReadBufferSize)
	readerProgressDecorator := component.NewReaderProgressDecorator(accessLogReader)
	ipProcessor, err := InitIPProcessor(processorConfig)
	if err != nil {
		return nil, err
	}
	filterAccessLogService := InitFilterAccessLogService(filterConfig)

	parser, err := mapper.NewAccessLogParser(pipelineConfig.LogFormat)
	if err != nil {
		return nil, err
	}

	return services.NewReadAccessLogService(readerProgressDecorator, filterAccessLogService, ipProcessor, pipelineConfig, parser), nil
}

// InitGeoReportService inits a services
func InitGeoReportService(maxmindGeoLibPath string, writeBufferSize int) *report.GeoReportService {
	geoProvider := geo.NewMaxMindGeoProvider(maxmindGeoLibPath)
	geoReportWriter := writer.NewGeoReportCsvWriter(writeBufferSize)

	return report.NewGeoReportService(geoProvider, geoReportWriter)
}

// InitBrowserReportService inits a services
func InitBrowserReportService(writeBufferSize int) *report.BrowserReportService {
	browserReportWriter := writer.NewBrowserReportCsvWriter(writeBufferSize)

	return report.NewBrowserReportService(browserReportWriter)
}

// InitRequestReportService inits a services
func InitRequestReportService(writeBufferSize int) *report.RequestReportService {
	requestReportWriter := writer.NewRequestReportCsvWriter(writeBufferSize)

	return report.NewRequestReportService(requestReportWriter)
}

// InitPaceReportService inits a services
func InitPaceReportService(writeBufferSize int) *report.PaceReportService {
	paceReportWriter := writer.NewPaceReportCsvWriter(writeBufferSize)

	return report.NewPaceReportService(paceReportWriter)
}

// InitJourneyReportService inits a services
func InitJourneyReportService(generalConfig config.GeneralConfig) *report.JourneyReportService {
	journeyReportWriter := writer.NewJourneyReportHTMLWriter()

	return report.NewJourneyReportService(generalConfig, journeyReportWriter)
}

// InitCustomReportService inits a services
func InitCustomReportService(writeBufferSize int) *report.CustomReportService {
	customReportWriter := writer.NewCustomReportCsvWriter(writeBufferSize)

	return report.NewCustomReportService(customReportWriter)
}
