package config

// PipelineConfig holds configuration for the concurrent processing pipeline.
type PipelineConfig struct {
	Workers        int
	ReadBufferSize int
}

// NewPipelineConfig creates a new PipelineConfig with the given parameters.
func NewPipelineConfig(
	workers int,
	readBufferSize int,
) PipelineConfig {
	return PipelineConfig{
		Workers:        workers,
		ReadBufferSize: readBufferSize,
	}
}
