package geodata

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"tango/internal/infrastructure/filesystem"
	"tango/internal/infrastructure/geo"
)

// InstallMaxmindLibUsecase usecase
type InstallMaxmindLibUsecase struct {
	homeDirResolver       *filesystem.HomeDirResolver
	maxmindGeoLibResolver *geo.MaxMindGeoLibResolver
}

// NewInstallMaxmindLibUsecase creates a new instance of InstallMaxMindLibraryUsecase
func NewInstallMaxmindLibUsecase(homeDirResolver *filesystem.HomeDirResolver, maxmindGeoLibResolver *geo.MaxMindGeoLibResolver) *InstallMaxmindLibUsecase {
	return &InstallMaxmindLibUsecase{
		homeDirResolver:       homeDirResolver,
		maxmindGeoLibResolver: maxmindGeoLibResolver,
	}
}

// Install installs MaxMind Geo library
func (u *InstallMaxmindLibUsecase) Install() {
	geoLibPath, geoLibExistError := u.maxmindGeoLibResolver.GetPath()

	if os.IsExist(geoLibExistError) {
		// MaxMind Geo Lib is already in place
		return
	}

	homeDirPath := u.homeDirResolver.GetPath()

	// installing of MaxMind Geo Lib
	geoLibArchivePath := filepath.Join(homeDirPath, "GeoLite2-City.tar.gz")

	u.downloadMaxMindArchive(geoLibArchivePath)
	u.untarMaxMindLib(geoLibArchivePath, geoLibPath)

	defer u.removeMaxMindArchive(geoLibArchivePath)
}

// downloadMaxMindArchive
func (u *InstallMaxmindLibUsecase) downloadMaxMindArchive(geoLibArchivePath string) {
	out, err := os.Create(geoLibArchivePath)
	defer out.Close()

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Get("https://geolite.maxmind.com/download/geoip/database/GeoLite2-City.tar.gz")

	if resp.StatusCode != http.StatusOK {
		log.Fatal(fmt.Errorf("Bad Response Status on Downloading Geo Lib: %s", resp.Status))
	}

	defer resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(out, resp.Body)

	if err != nil {
		log.Fatal(err)
	}
}

func (u *InstallMaxmindLibUsecase) untarMaxMindLib(geoLibArchivePath string, geoLibPath string) {
	libArchiveFile, err := os.Open(geoLibArchivePath)

	if err != nil {
		log.Fatal(err)
	}

	defer libArchiveFile.Close()

	// handle gzip
	libArchiveReader, err := gzip.NewReader(libArchiveFile)

	if err != nil {
		log.Fatal(err)
	}

	defer libArchiveReader.Close()

	tarBallReader := tar.NewReader(libArchiveReader)

	for {
		header, err := tarBallReader.Next()
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatal(err)
		}

		// get the individual filename and extract to the current directory
		tarredFileName := header.Name

		if header.Typeflag != tar.TypeReg {
			continue
		}

		if !strings.HasSuffix(tarredFileName, "GeoLite2-City.mmdb") {
			continue
		}

		// found geo lib in the tar ball archive

		geoLibWriter, err := os.Create(geoLibPath)

		if err != nil {
			log.Fatal(err)
		}

		io.Copy(geoLibWriter, tarBallReader)

		err = os.Chmod(geoLibPath, os.FileMode(header.Mode))

		if err != nil {
			log.Fatal(err)
		}

		geoLibWriter.Close()
	}
}

func (u *InstallMaxmindLibUsecase) removeMaxMindArchive(geoLibArchivePath string) {
	os.Remove(geoLibArchivePath)
}
