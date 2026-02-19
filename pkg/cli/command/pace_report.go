package command

import (
	"fmt"
	"tango/pkg/di"

	"github.com/urfave/cli"
)

// PaceReportCommand generates request pace reports
func PaceReportCommand(cliContext *cli.Context) error {
	reportConfig := di.InitReportConfig(cliContext)
	filterConfig := di.InitFilterConfig(cliContext)
	processorConfig := di.InitProcessorConfig(cliContext)
	pipelineConfig := di.InitPipelineConfig(cliContext)
	readAccessLogService, err := di.InitReadAccessLogService(processorConfig, filterConfig, pipelineConfig)
	if err != nil {
		return err
	}

	paceReportService := di.InitPaceReportService(pipelineConfig.WriteBufferSize)

	fmt.Println("💃 Tango is on the scene!")
	fmt.Println("💃 generating a request pace report...")

	records := readAccessLogService.Read(reportConfig.LogFile)
	paceReportService.GenerateReport(reportConfig.ReportFile, records)

	fmt.Println("🎉 request pace report has been generated")

	return nil
}
