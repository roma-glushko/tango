package config

// ReportConfig
type ReportConfig struct {
	LogFile    string
	ReportFile string
}

// NewReportConfig inits a config provider
func NewReportConfig(
	logFile string,
	reportFile string,
) ReportConfig {
	return ReportConfig{
		LogFile:    logFile,
		ReportFile: reportFile,
	}
}
