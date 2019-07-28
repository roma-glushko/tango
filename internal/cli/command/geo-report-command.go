package command

import (
	"fmt"
	"tango/internal/di"

	"github.com/urfave/cli"
)

//
func GeoReportCommand(c *cli.Context) error {
	readAccessLogUsecase := di.InitReadAccessLogUsecase()
	geoReportUsecase := di.InitGeoReportUsecase()

	fmt.Println("💃 Tango is on the scene!")
	fmt.Println("💃 started to generate a geo report...")
	fmt.Println("💃 reading access logs...")

	accessLogRecords := readAccessLogUsecase.Read("tmp/2019-07-08-transfer.log")

	fmt.Println("💃 saving the geo report...")

	geoReportUsecase.GenerateReport("output/geo.csv", accessLogRecords)

	fmt.Println("🎉 geo report has been generated")

	return nil
}
