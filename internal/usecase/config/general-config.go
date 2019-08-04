package config

//
type GeneralConfig struct {
	LogFile    string
	ReportFile string
	ConfigFile string
}

//
func NewGeneralConfig(
	logFile string,
	reportFile string,
	configFile string,
) GeneralConfig {
	return GeneralConfig{
		LogFile:    logFile,
		ReportFile: reportFile,
		ConfigFile: configFile,
	}
}
