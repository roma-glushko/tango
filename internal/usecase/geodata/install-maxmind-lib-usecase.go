package geodata

import (
	"fmt"
	"path/filepath"

	"github.com/maxmind/geoipupdate/v4/pkg/geoipupdate"
	"github.com/maxmind/geoipupdate/v4/pkg/geoipupdate/database"
)

// InstallMaxmindLibUsecase usecase
type InstallMaxmindLibUsecase struct {
}

// NewInstallMaxmindLibUsecase creates a new instance of InstallMaxMindLibraryUsecase
func NewInstallMaxmindLibUsecase() *InstallMaxmindLibUsecase {
	return &InstallMaxmindLibUsecase{}
}

// Install installs MaxMind Geo library
func (u *InstallMaxmindLibUsecase) Install(configFile string, dbDirectory string, verbose bool) {
	maxmindConfig, err := geoipupdate.NewConfig(configFile, dbDirectory, dbDirectory, verbose)

	if err != nil {
		fmt.Println("🚨 Error loading configuration file", err.Error())

		return
	}

	fmt.Println("🛠 Using MaxMind Config File: %s", configFile)
	fmt.Println("🛠 Using MaxMind DB Dir: %s", maxmindConfig.DatabaseDirectory)

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
