package cli

import (
	"fmt"
	"log"
	"tango/pkg/cli/command"
	"tango/pkg/cli/component"
	"time"

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
		// misc commands
		{
			Name:     "geo-lib",
			Aliases:  []string{"install-geo-lib", "get-geo-lib"},
			Category: "Misc",
			Usage:    "Install Geo Lib (from MaxMind).\nCredentials can be generated under https://www.maxmind.com/en/accounts/current/license-key page",
			Action:   command.InstallGeoLibCommand,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:     "update, u",
					Usage:    "Update/reinstall geo library",
					Required: false,
				},
				cli.StringFlag{
					Name:     "account-id, a",
					Usage:    "MaxMind AccountID",
					Required: false,
				},
				cli.StringFlag{
					Name:     "license-key, l",
					Usage:    "MaxMind LicenseKey",
					Required: false,
				},
			},
		},
		// report commands
		{
			Name:     "custom",
			Aliases:  []string{"custom-report"},
			Category: "Reports",
			Usage:    "Process Access Logs applying custom filtering",
			Action:   command.CustomReportCommand,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "log-file, l",
					Usage:    "Access log file to analyze",
					Required: true,
				},
				cli.StringFlag{
					Name:     "report-file, r",
					Usage:    "Output report file",
					Required: true,
				},
			},
		},
		{
			Name:     "geo",
			Aliases:  []string{"geo-report"},
			Category: "Reports",
			Usage:    "Generate Geo report from Access Logs",
			Action:   command.GeoReportCommand,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "log-file, l",
					Usage:    "Access log file to analyze",
					Required: true,
				},
				cli.StringFlag{
					Name:     "report-file, r",
					Usage:    "Output report file",
					Required: true,
				},
			},
		},
		{
			Name:     "browser",
			Aliases:  []string{"browser-report"},
			Category: "Reports",
			Usage:    "Generate Browser report from Access Logs",
			Action:   command.BrowserReportCommand,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "log-file, l",
					Usage:    "Access log file to analyze",
					Required: true,
				},
				cli.StringFlag{
					Name:     "report-file, r",
					Usage:    "Output report file",
					Required: true,
				},
			},
		},
		{
			Name:     "request",
			Aliases:  []string{"request-report"},
			Category: "Reports",
			Usage:    "Generate Request report from Access Logs",
			Action:   command.RequestReportCommand,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "log-file, l",
					Usage:    "Access log file to analyze",
					Required: true,
				},
				cli.StringFlag{
					Name:     "report-file, r",
					Usage:    "Output report file",
					Required: true,
				},
			},
		},
		{
			Name:     "journey",
			Aliases:  []string{"journey-report"},
			Category: "Reports",
			Usage:    "Generate a report based on visitor's journeys",
			Action:   command.JourneyReportCommand,
			Flags: []cli.Flag{
				cli.StringSliceFlag{
					Name:  "uri-filter, uf",
					Usage: "Remove URIs from visitor's journey, but consider these records during journey preparing",
				},
				cli.StringFlag{
					Name:     "log-file, l",
					Usage:    "Access log file to analyze",
					Required: true,
				},
				cli.StringFlag{
					Name:     "report-file, r",
					Usage:    "Output report file",
					Required: true,
				},
			},
		},
		{
			Name:     "pace",
			Aliases:  []string{"request-pace-report"},
			Category: "Reports",
			Usage:    "Generate request pace report from access logs",
			Action:   command.PaceReportCommand,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "log-file, l",
					Usage:    "Access log file to analyze",
					Required: true,
				},
				cli.StringFlag{
					Name:     "report-file, r",
					Usage:    "Output report file",
					Required: true,
				},
			},
		},
	}
}

// getTangoGlobalFlags returns all global filter flags
func getTangoGlobalFlags() []cli.Flag {
	return []cli.Flag{
		// general
		cli.StringFlag{Name: "config-file, c", Usage: "Configuration file for storing preset of filters", Value: component.DefaultConfigFileName},
		altsrc.NewStringFlag(cli.StringFlag{Name: "base-url", Usage: "Website Base Url"}),

		// filters
		altsrc.NewStringSliceFlag(cli.StringSliceFlag{Name: "asset-filter", Usage: "Filter js, css, image files from access logs"}),
		altsrc.NewStringSliceFlag(cli.StringSliceFlag{Name: "keep-time-filter", Usage: "Keep only specified time frame"}),
		altsrc.NewStringSliceFlag(cli.StringSliceFlag{Name: "uri-filter", Usage: "Filter URIs"}),
		altsrc.NewStringSliceFlag(cli.StringSliceFlag{Name: "keep-uri-filter", Usage: "Keep only specified URIs"}),
		altsrc.NewStringSliceFlag(cli.StringSliceFlag{Name: "ip-filter", Usage: "Filter IPs from access logs"}),
		altsrc.NewStringSliceFlag(cli.StringSliceFlag{Name: "keep-ip-filter", Usage: "Keep only specified IPs"}),
		altsrc.NewStringSliceFlag(cli.StringSliceFlag{Name: "ua-filter", Usage: "Filter specified user agents"}),
		altsrc.NewStringSliceFlag(cli.StringSliceFlag{Name: "keep-ua-filter", Usage: "Keep only specified user agents"}),

		// processors
		altsrc.NewStringSliceFlag(cli.StringSliceFlag{Name: "system-ips", Usage: "System IPs which are needed to filter like proxies"}),
	}
}

// NewTangoCli creates a main Tango CLI application
func NewTangoCli(version string, commit string) TangoCli {
	cliApp := cli.NewApp()
	today := time.Now()

	cliApp.Name = "Tango"
	cliApp.Usage = "Access Logs Analyzing Tool"
	cliApp.Version = fmt.Sprintf("%s (%s)", version, commit)
	cliApp.Copyright = fmt.Sprintf("(c) 2019-%d Roman Hlushko", today.Year())
	cliApp.Authors = []cli.Author{
		{
			Name:  "Roman Hlushko",
			Email: "roman.glushko.m@gmail.com",
		},
	}

	cliApp.Flags = getTangoGlobalFlags()
	cliApp.Commands = getTangoCommands()

	cliApp.Before = component.InitTangoConfigSourceWithContext(
		cliApp.Flags,
		component.NewTangoConfigYamlSourceFromFlagFunc("config-file"),
	)

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
