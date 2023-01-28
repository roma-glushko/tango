package filesystem

import (
	"log"
	"os"
	"path/filepath"
)

// HomeDirResolver retrieves Tango home dir path
type HomeDirResolver struct {
}

// NewHomeDirResolver creates a new instance of HomeDirResolver
func NewHomeDirResolver() *HomeDirResolver {
	return &HomeDirResolver{}
}

// GetPath provides path to Tango Home Dir
func (r *HomeDirResolver) GetPath() string {
	homeDirectory, err := os.UserHomeDir()

	if err != nil {
		log.Fatal(err)
	}

	tangoHomeDirectory := filepath.Join(homeDirectory, ".tango")

	// ensure that tango home dir is in place
	if _, err := os.Stat(tangoHomeDirectory); os.IsNotExist(err) {
		os.Mkdir(tangoHomeDirectory, os.ModePerm)
	}

	return tangoHomeDirectory
}
