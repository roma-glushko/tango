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
	homeDir, err := os.UserHomeDir()

	if err != nil {
		log.Fatal(err)
	}

	tangoDir := filepath.Join(homeDir, ".tango")

	// ensure that tango home dir is in place
	if _, err := os.Stat(tangoDir); os.IsNotExist(err) {
		err := os.Mkdir(tangoDir, os.ModePerm)

		if err != nil {
			log.Fatal(err)
		}
	}

	return tangoDir
}
