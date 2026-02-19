package command

import (
	"fmt"
	"tango/pkg/di"

	"github.com/urfave/cli"
)

// BrowserReportCommand handles the browser report CLI command.
func BrowserReportCommand(cliContext *cli.Context) error {
	reportConfig := di.InitReportConfig(cliContext)
	filterConfig := di.InitFilterConfig(cliContext)
	processorConfig := di.InitProcessorConfig(cliContext)
	pipelineConfig := di.InitPipelineConfig(cliContext)
	readAccessLogService, err := di.InitReadAccessLogService(processorConfig, filterConfig, pipelineConfig)
	if err != nil {
		return err
	}
	browserReportService := di.InitBrowserReportService(pipelineConfig.WriteBufferSize)

	fmt.Println("💃 Tango is on the scene!")
	fmt.Println("💃 generating a browser report...")

	records := readAccessLogService.Read(reportConfig.LogFile)
	browserReportService.GenerateReport(reportConfig.ReportFile, records)

	fmt.Println("🎉 browser report has been generated")

	return nil
}
