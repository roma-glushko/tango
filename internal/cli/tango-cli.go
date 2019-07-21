package cli

import (
	"log"
	"tango/internal/di"

	"github.com/urfave/cli"
)

type TangoCli struct {
	cliApp *cli.App
}

func getTangoCommands() []cli.Command {
	return []cli.Command{
		{
			Name:    "custom",
			Aliases: []string{},
			Usage:   "Process Access Logs applying custom filtering",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:    "geo",
			Aliases: []string{"geo-report"},
			Usage:   "Generate Geo report from Access Logs",
			Action: func(c *cli.Context) error {
				readAccessLogUsecase := di.InitReadAccessLogUsecase()

				readAccessLogUsecase.Read("tmp/2019-07-08-transfer.log")

				return nil
			},
		},
	}
}

// Create a main Tango CLI application
func NewTangoCli() TangoCli {
	cliApp := cli.NewApp()

	cliApp.Name = "Tango"
	cliApp.Usage = "Access Logs Analyzing Tool"
	cliApp.Version = "1.0.0-beta"

	cliApp.Commands = getTangoCommands()

	return TangoCli{
		cliApp: cliApp,
	}
}

// Execute Tango CLI Application
func (app *TangoCli) Run(arguments []string) {
	err := app.cliApp.Run(arguments)

	if err != nil {
		log.Fatal(err)
	}
}
