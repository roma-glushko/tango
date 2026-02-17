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
	fmt.Println("💃 started to generate a browser report...")
	fmt.Println("💃 reading access logs...")

	accessLogRecords := readAccessLogService.Read(reportConfig.LogFile)

	fmt.Println("💃 saving the browser report...")

	browserReportService.GenerateReport(reportConfig.ReportFile, accessLogRecords)

	fmt.Println("🎉 browser report has been generated")

	return nil
}
