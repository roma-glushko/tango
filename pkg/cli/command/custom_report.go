package command

import (
	"fmt"
	"tango/pkg/di"

	"github.com/urfave/cli"
)

// CustomReportCommand
func CustomReportCommand(cliContext *cli.Context) error {
	reportConfig := di.InitReportConfig(cliContext)
	filterConfig := di.InitFilterConfig(cliContext)
	processorConfig := di.InitProcessorConfig(cliContext)
	readAccessLogService, err := di.InitReadAccessLogService(processorConfig, filterConfig)
	if err != nil {
		return err
	}
	customReportService := di.InitCustomReportService()

	fmt.Println("💃 Tango is on the scene!")
	fmt.Println("💃 started to generate a custom report...")
	fmt.Println("💃 reading access logs...")

	accessLogRecords := readAccessLogService.Read(reportConfig.LogFile)

	fmt.Println("💃 saving the custom report...")

	customReportService.GenerateReport(reportConfig.ReportFile, accessLogRecords)

	fmt.Println("🎉 custom report has been generated")

	return nil
}
