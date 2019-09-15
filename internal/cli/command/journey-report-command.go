package command

import (
	"fmt"
	"tango/internal/di"

	"github.com/urfave/cli"
)

// JourneyReportCommand generates journey report for needed IPs
func JourneyReportCommand(cliContext *cli.Context) error {
	reportConfig := di.InitReportConfig(cliContext)
	generalConfig := di.InitGeneralConfig(cliContext)
	filterConfig := di.InitFilterConfig(cliContext)
	processorConfig := di.InitProcessorConfig(cliContext)
	readAccessLogUsecase := di.InitReadAccessLogUsecase(processorConfig, filterConfig)

	journeyReportUsecase := di.InitJourneyReportUsecase(generalConfig)

	fmt.Println("💃 Tango is on the scene!")
	fmt.Println("💃 started to generate a visitor's journey report...")
	fmt.Println("💃 reading access logs...")

	accessLogRecords := readAccessLogUsecase.Read(reportConfig.LogFile)

	fmt.Println("💃 saving visitor's journey report...")

	journeyReportUsecase.GenerateReport(reportConfig.ReportFile, accessLogRecords)

	fmt.Println("🎉 visitor's journey report has been generated")

	return nil
}
