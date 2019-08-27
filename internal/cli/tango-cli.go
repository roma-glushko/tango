package cli

import (
	"log"
	"tango/internal/cli/command"

	"github.com/urfave/cli"
	"github.com/urfave/cli/altsrc"
)

// TangoCli
type TangoCli struct {
	cliApp *cli.App
}

// getTangoCommands returns list of commands
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
		{
			Name:    "browser",
			Aliases: []string{"browser-report"},
			Usage:   "Generate Browser report from Access Logs",
			Action:  command.BrowserReportCommand,
		},
		{
			Name:    "request",
			Aliases: []string{"request-report"},
			Usage:   "Generate Request report from Access Logs",
			Action:  command.RequestReportCommand,
		},
		{
			Name:    "journey",
			Aliases: []string{"journey-report"},
			Usage:   "Generate a report based on visitor's journeys",
			Action:  command.JourneyReportCommand,
			Flags: []cli.Flag{
				cli.StringSliceFlag{
					Name:  "uri-filter, uf",
					Usage: "Remove URIs from visitor's journey, but consider these records during journey preparing",
				},
			},
		},
	}
}

// getTangoGlobalFlags returns all global filter flags
func getTangoGlobalFlags() []cli.Flag {
	return []cli.Flag{
		// general
		cli.StringFlag{Name: "log-file, l", Usage: "Access log file to analyze", Required: true},
		cli.StringFlag{Name: "report-file, r", Usage: "Output report file", Required: true},
		cli.StringFlag{Name: "config-file, c", Usage: "Configuration file for storing preset of filters", Value: ".tango.yaml"},

		// filters
		altsrc.NewStringSliceFlag(cli.StringSliceFlag{Name: "asset-filter", Usage: "Filter js, css, image files from access logs"}),
		altsrc.NewStringSliceFlag(cli.StringSliceFlag{Name: "keep-time-filter", Usage: "Keep only specified time frame"}),
		altsrc.NewStringSliceFlag(cli.StringSliceFlag{Name: "uri-filter", Usage: "Filter URIs"}),
		altsrc.NewStringSliceFlag(cli.StringSliceFlag{Name: "keep-uri-filter", Usage: "Keep only specified URIs"}),
		altsrc.NewStringSliceFlag(cli.StringSliceFlag{Name: "ip-filter", Usage: "Filter IPs from access logs"}),
		altsrc.NewStringSliceFlag(cli.StringSliceFlag{Name: "keep-ip-filter", Usage: "Keep only specified IPs"}),
		altsrc.NewStringSliceFlag(cli.StringSliceFlag{Name: "ua-filter", Usage: "Filter specified user agents"}),
		altsrc.NewStringSliceFlag(cli.StringSliceFlag{Name: "keep-ua-filter", Usage: "Keep only specified user agents"}),
		altsrc.NewStringFlag(cli.StringFlag{Name: "base-url", Usage: "Website Base Url"}),
		altsrc.NewStringSliceFlag(cli.StringSliceFlag{Name: "system-ips", Usage: "System IPs which are needed to filter like proxies"}),
	}
}

// NewTangoCli creates a main Tango CLI application
func NewTangoCli() TangoCli {
	cliApp := cli.NewApp()

	cliApp.Name = "Tango"
	cliApp.Usage = "Access Logs Analyzing Tool"
	cliApp.Version = "1.0.0-beta"

	cliApp.Flags = getTangoGlobalFlags()
	cliApp.Commands = getTangoCommands()

	cliApp.Before = altsrc.InitInputSourceWithContext(cliApp.Flags, altsrc.NewYamlSourceFromFlagFunc("config-file"))

	return TangoCli{
		cliApp: cliApp,
	}
}

// Run executes Tango CLI Application
func (app *TangoCli) Run(arguments []string) {
	err := app.cliApp.Run(arguments)

	if err != nil {
		log.Fatal(err)
	}
}
