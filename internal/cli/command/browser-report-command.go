package command

import (
	"fmt"
	"tango/internal/di"

	"github.com/urfave/cli"
)

//
func BrowserReportCommand(cliContext *cli.Context) error {
	generalConfig := di.InitGeneralConfig(cliContext)
	filterConfig := di.InitFilterConfig(cliContext)
	readAccessLogUsecase := di.InitReadAccessLogUsecase(filterConfig)
	browserReportUsecase := di.InitBrowserReportUsecase()

	fmt.Println("💃 Tango is on the scene!")
	fmt.Println("💃 started to generate a browser report...")
	fmt.Println("💃 reading access logs...")

	accessLogRecords := readAccessLogUsecase.Read(generalConfig.LogFile)

	fmt.Println("💃 saving the browser report...")

	browserReportUsecase.GenerateReport(generalConfig.ReportFile, accessLogRecords)

	fmt.Println("🎉 browser report has been generated")

	return nil
}
