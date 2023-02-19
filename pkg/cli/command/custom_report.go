package command

import (
	"fmt"
	"tango/pkg/di"
	"tango/pkg/services/mapper"

	"github.com/urfave/cli"
)

// CustomReportCommand
func CustomReportCommand(cliContext *cli.Context) error {
	reportConfig := di.InitReportConfig(cliContext)
	filterConfig := di.InitFilterConfig(cliContext)
	processorConfig := di.InitProcessorConfig(cliContext)

	logMapper := mapper.NewAccessLogMapper()
	readAccessLogService := di.InitReadAccessLogService(logMapper, processorConfig, filterConfig)
	customReportService := di.InitCustomReportService(logMapper)

	fmt.Println("💃 Tango is on the scene!")
	fmt.Println("💃 started to generate a custom report...")
	fmt.Println("💃 reading access logs...")

	logChan := readAccessLogService.Read(reportConfig.LogFile)

	fmt.Println("💃 saving the custom report...")

	customReportService.GenerateReport(reportConfig.ReportFile, logChan)

	fmt.Println("🎉 custom report has been generated")

	return nil
}
