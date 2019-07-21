package main

import (
	"os"
	"tango/internal/cli"
)

func main() {
	tangoCli := cli.NewTangoCli()

	tangoCli.Run(os.Args)
}
