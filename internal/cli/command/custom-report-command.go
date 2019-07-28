package command

import (
	"fmt"
	"tango/internal/di"

	"github.com/urfave/cli"
)

//
func CustomReportCommand(c *cli.Context) error {
	readAccessLogUsecase := di.InitReadAccessLogUsecase()
	customReportUsecase := di.InitCustomReportUsecase()

	fmt.Println("💃 Tango is on the scene!")
	fmt.Println("💃 started to generate a custom report...")
	fmt.Println("💃 reading access logs...")

	accessLogRecords := readAccessLogUsecase.Read("tmp/2019-07-08-transfer.log")

	fmt.Println("💃 saving the custom report...")

	customReportUsecase.GenerateReport("output/custom.csv", accessLogRecords)

	fmt.Println("🎉 custom report has been generated")

	return nil
}
