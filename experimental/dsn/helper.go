package main

import (
	"path/filepath"
	"runtime"
)

func fileInRuntimeDir(file string) string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename) + file
}
