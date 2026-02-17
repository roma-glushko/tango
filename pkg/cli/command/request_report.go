package command

import (
	"fmt"
	"tango/pkg/di"

	"github.com/urfave/cli"
)

// RequestReportCommand handles the request report CLI command.
func RequestReportCommand(cliContext *cli.Context) error {
	reportConfig := di.InitReportConfig(cliContext)
	filterConfig := di.InitFilterConfig(cliContext)
	processorConfig := di.InitProcessorConfig(cliContext)
	pipelineConfig := di.InitPipelineConfig(cliContext)
	readAccessLogService, err := di.InitReadAccessLogService(processorConfig, filterConfig, pipelineConfig)
	if err != nil {
		return err
	}
	requestReportService := di.InitRequestReportService(pipelineConfig.WriteBufferSize)

	fmt.Println("💃 Tango is on the scene!")
	fmt.Println("💃 generating a request report...")

	records := readAccessLogService.Read(reportConfig.LogFile)
	requestReportService.GenerateReport(reportConfig.ReportFile, records)

	fmt.Println("🎉 request report has been generated")

	return nil
}
