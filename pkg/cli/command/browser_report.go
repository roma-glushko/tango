package command

import (
	"fmt"
	"tango/pkg/di"
	"tango/pkg/services/mapper"

	"github.com/urfave/cli"
)

// BrowserReportCommand
func BrowserReportCommand(cliContext *cli.Context) error {
	reportConfig := di.InitReportConfig(cliContext)
	filterConfig := di.InitFilterConfig(cliContext)
	processorConfig := di.InitProcessorConfig(cliContext)

	logMapper := mapper.NewAccessLogMapper()
	readAccessLogService := di.InitReadAccessLogService(logMapper, processorConfig, filterConfig)
	browserReportService := di.InitBrowserReportService(logMapper)

	fmt.Println("💃 Tango is on the scene!")
	fmt.Println("💃 started to generate a browser report...")
	fmt.Println("💃 reading access logs...")

	logChan := readAccessLogService.Read(reportConfig.LogFile)

	fmt.Println("💃 saving the browser report...")

	browserReportService.GenerateReport(reportConfig.ReportFile, logChan)

	fmt.Println("🎉 browser report has been generated")

	return nil
}
