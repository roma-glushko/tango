package cli

import (
	"log"
	"tango/internal/cli/command"

	"github.com/urfave/cli"
)

//
type TangoCli struct {
	cliApp *cli.App
}

func getTangoCommands() []cli.Command {
	return []cli.Command{
		{
			Name:    "custom",
			Aliases: []string{"custom-report"},
			Usage:   "Process Access Logs applying custom filtering",
			Action:  command.CustomReportCommand,
		},
		{
			Name:    "geo",
			Aliases: []string{"geo-report"},
			Usage:   "Generate Geo report from Access Logs",
			Action:  command.GeoReportCommand,
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
