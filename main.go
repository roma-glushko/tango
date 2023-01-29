package main

import (
	"os"
	"tango/pkg/cli"
)

var (
	version = "dev"
	commit  = "none"
)

func main() {
	tangoCli := cli.NewTangoCli(version, commit)

	tangoCli.Run(os.Args)
}
