package geodata

import (
	"fmt"
	"path/filepath"

	"github.com/maxmind/geoipupdate/v4/pkg/geoipupdate"
	"github.com/maxmind/geoipupdate/v4/pkg/geoipupdate/database"
)

// InstallMaxmindLibService services
type InstallMaxmindLibService struct {
}

// NewInstallMaxmindLibService creates a new instance of InstallMaxMindLibraryService
func NewInstallMaxmindLibService() *InstallMaxmindLibService {
	return &InstallMaxmindLibService{}
}

// Install installs MaxMind Geo library
func (u *InstallMaxmindLibService) Install(configFile string, dbDirectory string, verbose bool) {
	maxmindConfig, err := geoipupdate.NewConfig(configFile, dbDirectory, dbDirectory, verbose)

	if err != nil {
		fmt.Println("🚨 Error loading configuration file", err.Error())

		return
	}

	fmt.Printf("🛠 Using MaxMind Config File: %s\n", configFile)
	fmt.Printf("🛠 Using MaxMind DB Dir: %s\n", maxmindConfig.DatabaseDirectory)

	client := geoipupdate.NewClient(maxmindConfig)

	dbReader := database.NewHTTPDatabaseReader(client, maxmindConfig)

	for _, editionID := range maxmindConfig.EditionIDs {
		filename, err := geoipupdate.GetFilename(maxmindConfig, editionID, client)

		if err != nil {
			fmt.Println("🚨 Error retrieving filename", err.Error())

			return
		}

		filePath := filepath.Join(maxmindConfig.DatabaseDirectory, filename)
		dbWriter, err := database.NewLocalFileDatabaseWriter(filePath, maxmindConfig.LockFile, maxmindConfig.Verbose)

		if err != nil {
			fmt.Println("🚨 Error creating MaxMind db writer", err.Error())

			return
		}

		if err := dbReader.Get(dbWriter, editionID); err != nil {
			fmt.Println("🚨", err.Error())

			return
		}
	}
}
