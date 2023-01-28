package config

//
type GeneralConfig struct {
	ConfigFile string
	BaseURL    string
}

// NewGeneralConfig inits a config provider
func NewGeneralConfig(
	configFile string,
	baseURL string,
) GeneralConfig {
	return GeneralConfig{
		ConfigFile: configFile,
		BaseURL:    baseURL,
	}
}
