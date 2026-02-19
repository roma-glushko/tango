package command

import (
	"fmt"
	"os"
	"tango/pkg/di"

	"github.com/urfave/cli"
)

// GeoReportCommand handles the geo report CLI command.
func GeoReportCommand(cliContext *cli.Context) error {
	reportConfig := di.InitReportConfig(cliContext)
	filterConfig := di.InitFilterConfig(cliContext)
	processorConfig := di.InitProcessorConfig(cliContext)
	pipelineConfig := di.InitPipelineConfig(cliContext)
	readAccessLogService, err := di.InitReadAccessLogService(processorConfig, filterConfig, pipelineConfig)
	if err != nil {
		return err
	}
	geoLibResolver := di.InitMaxmindGeoLibResolver()

	fmt.Println("💃 Tango is on the scene!")

	geoLibPath, err := geoLibResolver.GetPath()

	// ensure that geo library is in place
	if os.IsNotExist(err) {
		fmt.Println("🚨 Cannot perform geo reports without MaxMind geo database installed")
		fmt.Println("🚨 Please run 'tango geo-lib -h' to get more info about installation")

		return nil
	}

	geoReportService := di.InitGeoReportService(geoLibPath, pipelineConfig.WriteBufferSize)

	fmt.Println("💃 generating a geo report...")

	records := readAccessLogService.Read(reportConfig.LogFile)
	geoReportService.GenerateReport(reportConfig.ReportFile, records)

	fmt.Println("🎉 geo report has been generated")

	return nil
}
