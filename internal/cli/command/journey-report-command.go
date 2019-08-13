package command

import (
	"fmt"
	"tango/internal/di"

	"github.com/urfave/cli"
)

//
func JourneyReportCommand(cliContext *cli.Context) error {
	generalConfig := di.InitGeneralConfig(cliContext)
	filterConfig := di.InitFilterConfig(cliContext)
	processorConfig := di.InitProcessorConfig(cliContext)
	readAccessLogUsecase := di.InitReadAccessLogUsecase(processorConfig, filterConfig)

	requestReportUsecase := di.InitRequestReportUsecase() // todo: replace by actual command

	fmt.Println("💃 Tango is on the scene!")
	fmt.Println("💃 started to generate a visitor's journey report...")
	fmt.Println("💃 reading access logs...")

	accessLogRecords := readAccessLogUsecase.Read(generalConfig.LogFile)

	fmt.Println("💃 saving the visitor's journey report...")

	requestReportUsecase.GenerateReport(generalConfig.ReportFile, accessLogRecords)

	fmt.Println("🎉 visitor's journey report has been generated")

	return nil
}
