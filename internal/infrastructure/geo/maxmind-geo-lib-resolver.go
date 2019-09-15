package geo

import (
	"os"
	"path/filepath"
	"tango/internal/infrastructure/filesystem"
)

// MaxMindGeoLibResolver retrieves path to the MaxMind Geo Library
type MaxMindGeoLibResolver struct {
	homeDirResolver *filesystem.HomeDirResolver
}

// NewMaxMindGeoLibResolver creates a new instance of MaxMindGeoLibResolver
func NewMaxMindGeoLibResolver(homeDirResolver *filesystem.HomeDirResolver) *MaxMindGeoLibResolver {
	return &MaxMindGeoLibResolver{
		homeDirResolver: homeDirResolver,
	}
}

// GetPath provides path to MaxMind Geo Library
// there is only one possible path where the lib file can be found: $HOME/.tango/GeoLite2-City.mmdb
// $HOME path should be different on diff OS
func (r *MaxMindGeoLibResolver) GetPath() (string, error) {
	homeDirectory := r.homeDirResolver.GetPath()

	maxmindGeoLibPath := filepath.Join(homeDirectory, "GeoLite2-City.mmdb")
	_, maxmindGeoLibExistanceError := os.Stat(maxmindGeoLibPath)

	return maxmindGeoLibPath, maxmindGeoLibExistanceError
}
