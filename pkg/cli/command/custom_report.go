package command

import (
	"fmt"
	"tango/pkg/di"

	"github.com/urfave/cli"
)

// CustomReportCommand handles the custom report CLI command.
func CustomReportCommand(cliContext *cli.Context) error {
	reportConfig := di.InitReportConfig(cliContext)
	filterConfig := di.InitFilterConfig(cliContext)
	processorConfig := di.InitProcessorConfig(cliContext)
	pipelineConfig := di.InitPipelineConfig(cliContext)
	readAccessLogService, err := di.InitReadAccessLogService(processorConfig, filterConfig, pipelineConfig)
	if err != nil {
		return err
	}
	customReportService := di.InitCustomReportService(pipelineConfig.WriteBufferSize)

	fmt.Println("💃 Tango is on the scene!")
	fmt.Println("💃 generating a custom report...")

	records := readAccessLogService.Read(reportConfig.LogFile)
	customReportService.GenerateReport(reportConfig.ReportFile, records)

	fmt.Println("🎉 custom report has been generated")

	return nil
}
