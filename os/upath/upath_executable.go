package upath

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// ExecutableFolder
func Executable() string {
	path, err := os.Executable()
	if err != nil {
		return ""
	}
	return filepath.Dir(filepath.Clean(path))
}

// ExecutableSourceRoot
func ExecutableSourceRoot() string {
	_, currentPath, _, _ := runtime.Caller(0)
	return strings.Split(path.Dir(currentPath), "os")[0]
}
