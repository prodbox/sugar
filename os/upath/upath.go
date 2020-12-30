package upath

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

// ExecutableFolder 可执行文件路径
func Executable() string {
	path, err := os.Executable()
	if err != nil {
		return ""
	}
	return filepath.Dir(filepath.Clean(path))
}

// ExecutableSourceFolder 可执行文件源码路径
func ExecutableSourceRoot() string {
	fmt.Println(filepath.Abs(filepath.Dir(os.Args[0])))
	_, currentPath, _, _ := runtime.Caller(0)
	return path.Dir(currentPath)
}
