package config

//
type ProcessorConfig struct {
	SystemIpList []string
}

//
func NewProcessorConfig(
	systemIpList []string,
) ProcessorConfig {
	return ProcessorConfig{
		SystemIpList: systemIpList,
	}
}
