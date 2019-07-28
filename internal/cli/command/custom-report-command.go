package command

import (
	"tango/internal/di"

	"github.com/urfave/cli"
)

//
func CustomReportCommand(c *cli.Context) error {
	readAccessLogUsecase := di.InitReadAccessLogUsecase()
	customReportUsecase := di.InitCustomReportUsecase()

	accessLogRecords := readAccessLogUsecase.Read("tmp/2019-07-08-transfer.log")

	customReportUsecase.GenerateReport("output/custom.csv", accessLogRecords)

	return nil
}
