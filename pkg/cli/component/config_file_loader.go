package component

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
	"github.com/urfave/cli/altsrc"
)

// DefaultConfigFileName defines config file
var DefaultConfigFileName = ".tango.yaml"

// NewTangoConfigYamlSourceFromFlagFunc loads tango default config file if it's created or don't load anything it
// https://github.com/urfave/cli/issues/706
func NewTangoConfigYamlSourceFromFlagFunc(flagFileName string) func(context *cli.Context) (altsrc.InputSourceContext, error) {
	return func(context *cli.Context) (altsrc.InputSourceContext, error) {
		filePath := context.String(flagFileName)
		_, configFileError := os.Stat(filePath)

		if filePath == DefaultConfigFileName && os.IsNotExist(configFileError) {
			return nil, nil
		}

		return altsrc.NewYamlSourceFromFile(filePath)
	}
}

// InitTangoConfigSourceWithContext inits a fix for well-know issue
// https://github.com/urfave/cli/issues/706
func InitTangoConfigSourceWithContext(flags []cli.Flag, createInputSource func(context *cli.Context) (altsrc.InputSourceContext, error)) cli.BeforeFunc {
	return func(context *cli.Context) error {
		inputSource, err := createInputSource(context)

		if inputSource == nil {
			return nil
		}

		if err != nil {
			return fmt.Errorf("Unable to create input source with context: inner error: \n'%v'", err.Error())
		}

		return altsrc.ApplyInputSourceValues(context, inputSource, flags)
	}
}
