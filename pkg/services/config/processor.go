package config

// ProcessorConfig holds configuration for access log processors.
type ProcessorConfig struct {
	SystemIPList []string
}

// NewProcessorConfig creates a new ProcessorConfig with the given parameters.
func NewProcessorConfig(
	systemIPList []string,
) ProcessorConfig {
	return ProcessorConfig{
		SystemIPList: systemIPList,
	}
}
