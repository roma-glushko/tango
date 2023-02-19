package command

import (
	"fmt"
	"tango/pkg/di"
	"tango/pkg/services/mapper"

	"github.com/urfave/cli"
)

// RequestReportCommand
func RequestReportCommand(cliContext *cli.Context) error {
	reportConfig := di.InitReportConfig(cliContext)
	filterConfig := di.InitFilterConfig(cliContext)
	processorConfig := di.InitProcessorConfig(cliContext)

	logMapper := mapper.NewAccessLogMapper()
	readAccessLogService := di.InitReadAccessLogService(logMapper, processorConfig, filterConfig)
	requestReportService := di.InitRequestReportService(logMapper)

	fmt.Println("💃 Tango is on the scene!")
	fmt.Println("💃 started to generate a request report...")
	fmt.Println("💃 reading access logs...")

	logChan := readAccessLogService.Read(reportConfig.LogFile)
	requestReportService.GenerateReport(reportConfig.ReportFile, logChan)

	fmt.Println("🎉 request report has been generated")

	return nil
}
