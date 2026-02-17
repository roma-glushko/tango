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
	fmt.Println("💃 started to generate a request report...")
	fmt.Println("💃 reading access logs...")

	accessLogRecords := readAccessLogService.Read(reportConfig.LogFile)

	fmt.Println("💃 saving the request report...")

	requestReportService.GenerateReport(reportConfig.ReportFile, accessLogRecords)

	fmt.Println("🎉 request report has been generated")

	return nil
}
