package command

import (
	"fmt"
	"os"
	"tango/internal/di"

	"github.com/urfave/cli"
)

//
func GeoReportCommand(cliContext *cli.Context) error {
	generalConfig := di.InitGeneralConfig(cliContext)
	filterConfig := di.InitFilterConfig(cliContext)
	processorConfig := di.InitProcessorConfig(cliContext)
	readAccessLogUsecase := di.InitReadAccessLogUsecase(processorConfig, filterConfig)
	geoLibResolver := di.InitMaxmindGeoLibResolver()

	fmt.Println("💃 Tango is on the scene!")

	geoLibPath, err := geoLibResolver.GetPath()

	// ensure that geo library is in place
	if os.IsNotExist(err) {
		installMaxmindLibUsecase := di.InitInstallMaxmindLibUsecase()

		fmt.Println("💃 installing geo lib...")
		installMaxmindLibUsecase.Install()
		fmt.Println("🎉 geo lib has been installed")
	}

	geoReportUsecase := di.InitGeoReportUsecase(geoLibPath)

	fmt.Println("💃 started to generate a geo report...")
	fmt.Println("💃 reading access logs...")

	accessLogRecords := readAccessLogUsecase.Read(generalConfig.LogFile)

	fmt.Println("💃 saving the geo report...")

	geoReportUsecase.GenerateReport(generalConfig.ReportFile, accessLogRecords)

	fmt.Println("🎉 geo report has been generated")

	return nil
}
