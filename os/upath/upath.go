package upath

import (
	"os"
	"path/filepath"
)

// ExecutableFolder
func ExecutableFolder() string {
	path, err := os.Executable()
	if err != nil {
		return ""
	}
	return filepath.Dir(filepath.Clean(path))
}
