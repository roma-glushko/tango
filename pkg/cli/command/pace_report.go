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
	readAccessLogService, err := di.InitReadAccessLogService(processorConfig, filterConfig)
	if err != nil {
		return err
	}

	paceReportService := di.InitPaceReportService()

	fmt.Println("💃 Tango is on the scene!")
	fmt.Println("💃 started to generate a request pace report...")
	fmt.Println("💃 reading access logs...")

	accessLogRecords := readAccessLogService.Read(reportConfig.LogFile)

	fmt.Println("💃 saving the request pace report...")

	paceReportService.GenerateReport(reportConfig.ReportFile, accessLogRecords)

	fmt.Println("🎉 request pace report has been generated")

	return nil
}
