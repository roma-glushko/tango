package command

import (
	"fmt"
	"tango/pkg/di"

	"github.com/urfave/cli"
)

// JourneyReportCommand generates journey report for needed IPs
func JourneyReportCommand(cliContext *cli.Context) error {
	reportConfig := di.InitReportConfig(cliContext)
	generalConfig := di.InitGeneralConfig(cliContext)
	filterConfig := di.InitFilterConfig(cliContext)
	processorConfig := di.InitProcessorConfig(cliContext)
	pipelineConfig := di.InitPipelineConfig(cliContext)
	readAccessLogService, err := di.InitReadAccessLogService(processorConfig, filterConfig, pipelineConfig)
	if err != nil {
		return err
	}

	journeyReportService := di.InitJourneyReportService(generalConfig)

	fmt.Println("💃 Tango is on the scene!")
	fmt.Println("💃 started to generate a visitor's journey report...")
	fmt.Println("💃 reading access logs...")

	accessLogRecords := readAccessLogService.Read(reportConfig.LogFile)

	fmt.Println("💃 saving visitor's journey report...")

	journeyReportService.GenerateReport(reportConfig.ReportFile, accessLogRecords)

	fmt.Println("🎉 visitor's journey report has been generated")

	return nil
}
