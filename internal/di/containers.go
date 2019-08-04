package di

import (
	"tango/internal/cli/component"
	"tango/internal/infrastructure/reader"
	"tango/internal/infrastructure/writer"
	"tango/internal/usecase"
	"tango/internal/usecase/filter"
	"tango/internal/usecase/report"
)

//
func InitFilterAccessLogUsecase() usecase.FilterAccessLogUsecase {
	filters := []usecase.AccessLogFilter{
		filter.NewIPFilter(),
		filter.NewTimeFilter(),
		filter.NewUrlFilter(),
		filter.NewWebAssetFilter(),
		filter.NewUserAgentFilter(),
	}

	return usecase.NewFilterAccessLogUsecase(filters)
}

//
func InitReadAccessLogUsecase() *usecase.ReadAccessLogUsecase {
	accessLogReader := reader.NewAccessLogReader()
	readerProgressDecorator := component.NewReaderProgressDecorator(accessLogReader)
	filterAccessLogUsecase := InitFilterAccessLogUsecase()

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
