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
	fmt.Println("💃 generating a visitor's journey report...")

	records := readAccessLogService.Read(reportConfig.LogFile)
	journeyReportService.GenerateReport(reportConfig.ReportFile, records)

	fmt.Println("🎉 visitor's journey report has been generated")

	return nil
}
