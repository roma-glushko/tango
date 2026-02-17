package config

// PipelineConfig holds configuration for the concurrent processing pipeline.
type PipelineConfig struct {
	Workers        int
	ReadBufferSize int
	LogFormat      string
}

// NewPipelineConfig creates a new PipelineConfig with the given parameters.
func NewPipelineConfig(
	workers int,
	readBufferSize int,
	logFormat string,
) PipelineConfig {
	return PipelineConfig{
		Workers:        workers,
		ReadBufferSize: readBufferSize,
		LogFormat:      logFormat,
	}
}
