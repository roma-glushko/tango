package command

import (
	"fmt"
	"tango/internal/di"

	"github.com/urfave/cli"
)

// PaceReportCommand generates request pace reports
func PaceReportCommand(cliContext *cli.Context) error {
	reportConfig := di.InitReportConfig(cliContext)
	filterConfig := di.InitFilterConfig(cliContext)
	processorConfig := di.InitProcessorConfig(cliContext)
	readAccessLogUsecase := di.InitReadAccessLogUsecase(processorConfig, filterConfig)

	paceReportUsecase := di.InitPaceReportUsecase()

	fmt.Println("💃 Tango is on the scene!")
	fmt.Println("💃 started to generate a request pace report...")
	fmt.Println("💃 reading access logs...")

	accessLogRecords := readAccessLogUsecase.Read(reportConfig.LogFile)

	fmt.Println("💃 saving the request pace report...")

	paceReportUsecase.GenerateReport(reportConfig.ReportFile, accessLogRecords)

	fmt.Println("🎉 request pace report has been generated")

	return nil
}
