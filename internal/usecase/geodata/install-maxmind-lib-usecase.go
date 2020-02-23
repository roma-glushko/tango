package geodata

import (
	"fmt"
	"log"
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
		fmt.Errorf("error loading configuration file", err)
	}

	if maxmindConfig.Verbose {
		log.Printf("Using config file %s", configFile)
		log.Printf("Using database directory %s", maxmindConfig.DatabaseDirectory)
	}

	client := geoipupdate.NewClient(maxmindConfig)

	dbReader := database.NewHTTPDatabaseReader(client, maxmindConfig)

	for _, editionID := range maxmindConfig.EditionIDs {
		filename, err := geoipupdate.GetFilename(maxmindConfig, editionID, client)

		if err != nil {
			//return errors.Wrap(err, "error retrieving filename")
		}

		filePath := filepath.Join(maxmindConfig.DatabaseDirectory, filename)
		dbWriter, err := database.NewLocalFileDatabaseWriter(filePath, maxmindConfig.LockFile, maxmindConfig.Verbose)

		if err != nil {
			//return errors.Wrap(err, "error creating database writer")
		}

		if err := dbReader.Get(dbWriter, editionID); err != nil {
			//return err
		}
	}
}
