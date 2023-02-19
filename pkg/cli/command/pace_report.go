package command

import (
	"fmt"
	"tango/pkg/di"
	"tango/pkg/services/mapper"

	"github.com/urfave/cli"
)

// PaceReportCommand generates request pace reports
func PaceReportCommand(cliContext *cli.Context) error {
	reportConfig := di.InitReportConfig(cliContext)
	filterConfig := di.InitFilterConfig(cliContext)
	processorConfig := di.InitProcessorConfig(cliContext)

	logMapper := mapper.NewAccessLogMapper()
	readAccessLogService := di.InitReadAccessLogService(logMapper, processorConfig, filterConfig)
	paceReportService := di.InitPaceReportService(logMapper)

	fmt.Println("💃 Tango is on the scene!")
	fmt.Println("💃 started to generate a request pace report...")
	fmt.Println("💃 reading access logs...")

	logChan := readAccessLogService.Read(reportConfig.LogFile)

	fmt.Println("💃 saving the request pace report...")

	paceReportService.GenerateReport(reportConfig.ReportFile, logChan)

	fmt.Println("🎉 request pace report has been generated")

	return nil
}
