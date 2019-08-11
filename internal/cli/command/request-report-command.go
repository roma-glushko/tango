package command

import (
	"fmt"
	"tango/internal/di"

	"github.com/urfave/cli"
)

//
func RequestReportCommand(cliContext *cli.Context) error {
	generalConfig := di.InitGeneralConfig(cliContext)
	filterConfig := di.InitFilterConfig(cliContext)
	processorConfig := di.InitProcessorConfig(cliContext)
	readAccessLogUsecase := di.InitReadAccessLogUsecase(processorConfig, filterConfig)
	requestReportUsecase := di.InitRequestReportUsecase()

	fmt.Println("💃 Tango is on the scene!")
	fmt.Println("💃 started to generate a request report...")
	fmt.Println("💃 reading access logs...")

	accessLogRecords := readAccessLogUsecase.Read(generalConfig.LogFile)

	fmt.Println("💃 saving the request report...")

	requestReportUsecase.GenerateReport(generalConfig.ReportFile, accessLogRecords)

	fmt.Println("🎉 request report has been generated")

	return nil
}
