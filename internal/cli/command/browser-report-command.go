package command

import (
	"fmt"
	"tango/internal/di"

	"github.com/urfave/cli"
)

//
func BrowserReportCommand(cliContext *cli.Context) error {
	reportConfig := di.InitReportConfig(cliContext)
	filterConfig := di.InitFilterConfig(cliContext)
	processorConfig := di.InitProcessorConfig(cliContext)
	readAccessLogUsecase := di.InitReadAccessLogUsecase(processorConfig, filterConfig)
	browserReportUsecase := di.InitBrowserReportUsecase()

	fmt.Println("💃 Tango is on the scene!")
	fmt.Println("💃 started to generate a browser report...")
	fmt.Println("💃 reading access logs...")

	accessLogRecords := readAccessLogUsecase.Read(reportConfig.LogFile)

	fmt.Println("💃 saving the browser report...")

	browserReportUsecase.GenerateReport(reportConfig.ReportFile, accessLogRecords)

	fmt.Println("🎉 browser report has been generated")

	return nil
}
