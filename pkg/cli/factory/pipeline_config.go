package factory

import (
	"runtime"
	"tango/pkg/services/config"
	"tango/pkg/services/mapper"

	"github.com/urfave/cli"
)

const (
	// DefaultReadBufferSize is the default buffer size for reading log files (256KB).
	DefaultReadBufferSize = 256 * 1024
	// DefaultWriteBufferSize is the default buffer size for writing report files (256KB).
	DefaultWriteBufferSize = 256 * 1024
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

	writeBufferSize := cliContext.GlobalInt("write-buffer-size")
	if writeBufferSize <= 0 {
		writeBufferSize = DefaultWriteBufferSize
	}

	logFormat := cliContext.GlobalString("log-format")
	if logFormat == "" {
		logFormat = mapper.DefaultLogFormat
	}

	return config.NewPipelineConfig(
		workers,
		readBufferSize,
		writeBufferSize,
		logFormat,
	)
}
