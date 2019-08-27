package config

//
type GeneralConfig struct {
	LogFile    string
	ReportFile string
	ConfigFile string
	BaseURL    string
}

// NewGeneralConfig inits a config provider
func NewGeneralConfig(
	logFile string,
	reportFile string,
	configFile string,
	baseURL string,
) GeneralConfig {
	return GeneralConfig{
		LogFile:    logFile,
		ReportFile: reportFile,
		ConfigFile: configFile,
		BaseURL:    baseURL,
	}
}
