package command

import (
	"fmt"
	"os"
	"tango/pkg/di"
	"tango/pkg/services/mapper"

	"github.com/urfave/cli"
)

// GeoReportCommand
func GeoReportCommand(cliContext *cli.Context) error {
	reportConfig := di.InitReportConfig(cliContext)
	filterConfig := di.InitFilterConfig(cliContext)
	processorConfig := di.InitProcessorConfig(cliContext)

	logMapper := mapper.NewAccessLogMapper()
	readAccessLogService := di.InitReadAccessLogService(logMapper, processorConfig, filterConfig)
	geoLibResolver := di.InitMaxmindGeoLibResolver()

	fmt.Println("💃 Tango is on the scene!")

	geoLibPath, err := geoLibResolver.GetPath()

	// ensure that geo library is in place
	if os.IsNotExist(err) {
		fmt.Println("🚨 Cannot perform geo reports without MaxMind geo database installed")
		fmt.Println("🚨 Please run 'tango geo-lib -h' to get more info about installation")

		return nil
	}

	geoReportService := di.InitGeoReportService(logMapper, geoLibPath)

	fmt.Println("💃 started to generate a geo report...")
	fmt.Println("💃 reading access logs...")

	logChan := readAccessLogService.Read(reportConfig.LogFile)

	fmt.Println("💃 saving the geo report...")

	geoReportService.GenerateReport(reportConfig.ReportFile, logChan)

	fmt.Println("🎉 geo report has been generated")

	return nil
}
