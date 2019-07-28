package command

import (
	"tango/internal/di"

	"github.com/urfave/cli"
)

//
func GeoReportCommand(c *cli.Context) error {
	readAccessLogUsecase := di.InitReadAccessLogUsecase()
	geoReportUsecase := di.InitGeoReportUsecase()

	accessLogRecords := readAccessLogUsecase.Read("tmp/2019-07-08-transfer.log")

	geoReportUsecase.GenerateReport("output/geo.csv", accessLogRecords)

	return nil
}
