package geo

import (
	"os"
	"path/filepath"
	"tango/pkg/adapters/filesystem"
)

// MaxMindConfResolver retrieves path to the MaxMind Geo Confi file
type MaxMindConfResolver struct {
	homeDirResolver *filesystem.HomeDirResolver
}

// NewMaxMindConfResolver creates a new instance of MaxMindConfResolver
func NewMaxMindConfResolver(homeDirResolver *filesystem.HomeDirResolver) *MaxMindConfResolver {
	return &MaxMindConfResolver{
		homeDirResolver: homeDirResolver,
	}
}

// GetPath provides path to MaxMind Geo Library
// there is only one possible path where the lib file can be found: $HOME/.tango/maxmind.conf
// $HOME path should be different on diff OS
func (r *MaxMindConfResolver) GetPath() (string, error) {
	homeDirectory := r.homeDirResolver.GetPath()

	maxmindConfPath := filepath.Join(homeDirectory, "maxmind.conf")
	_, maxmindConfExistError := os.Stat(maxmindConfPath)

	return maxmindConfPath, maxmindConfExistError
}
