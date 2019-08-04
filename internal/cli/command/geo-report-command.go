package command

import (
	"fmt"
	"tango/internal/di"

	"github.com/urfave/cli"
)

//
func GeoReportCommand(cliContext *cli.Context) error {
	generalConfig := di.InitGeneralConfig(cliContext)
	filterConfig := di.InitFilterConfig(cliContext)
	readAccessLogUsecase := di.InitReadAccessLogUsecase(filterConfig)
	geoReportUsecase := di.InitGeoReportUsecase()

	fmt.Println("💃 Tango is on the scene!")
	fmt.Println("💃 started to generate a geo report...")
	fmt.Println("💃 reading access logs...")

	accessLogRecords := readAccessLogUsecase.Read(generalConfig.LogFile)

	fmt.Println("💃 saving the geo report...")

	geoReportUsecase.GenerateReport(generalConfig.ReportFile, accessLogRecords)

	fmt.Println("🎉 geo report has been generated")

	return nil
}
