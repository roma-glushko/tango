package geodata

import (
	"fmt"
	"os"
)

// GenerateMaxmindConfService services
type GenerateMaxmindConfService struct {
}

// NewGenerateMaxmindConfService creates a new instance of InstallMaxMindLibraryService
func NewGenerateMaxmindConfService() *GenerateMaxmindConfService {
	return &GenerateMaxmindConfService{}
}

// Generate MaxMind Conf File
func (u *GenerateMaxmindConfService) Generate(confPath string, accountID string, licenseKey string) error {
	confFile, err := os.Create(confPath)

	if err != nil {
		return err
	}

	defer confFile.Close()

	confFile.WriteString("# This MaxMind config file was generated by Tango \n")
	confFile.WriteString(fmt.Sprintf("AccountID %s\n", accountID))
	confFile.WriteString(fmt.Sprintf("LicenseKey %s\n", licenseKey))
	confFile.WriteString(fmt.Sprintf("EditionIDs %s\n", u.GetEditionIDs()))

	confFile.Sync()

	return nil
}

// GetEditionIDs retrieves MaxMind Product IDs which are used by Tango
func (u *GenerateMaxmindConfService) GetEditionIDs() string {
	return "GeoLite2-City"
}