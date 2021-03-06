// +build !windows

package upath

import (
	"os"
	"syscall"
)

// IsExist
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// IsDir
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// Mkdir
func Mkdir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// IsWritable
func IsWritable(path string) bool {
	return syscall.Access(path, syscall.O_RDWR) == nil
}
