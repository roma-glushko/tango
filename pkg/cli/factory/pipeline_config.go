package factory

import (
	"runtime"
	"tango/pkg/services/config"

	"github.com/urfave/cli"
)

const (
	// DefaultReadBufferSize is the default buffer size for reading log files (64KB).
	DefaultReadBufferSize = 64 * 1024
)

// NewPipelineConfig creates a PipelineConfig from CLI context flags.
func NewPipelineConfig(cliContext *cli.Context) config.PipelineConfig {
	workers := cliContext.GlobalInt("workers")
	if workers <= 0 {
		workers = runtime.NumCPU()
	}

	readBufferSize := cliContext.GlobalInt("read-buffer-size")
	if readBufferSize <= 0 {
		readBufferSize = DefaultReadBufferSize
	}

	return config.NewPipelineConfig(
		workers,
		readBufferSize,
	)
}
