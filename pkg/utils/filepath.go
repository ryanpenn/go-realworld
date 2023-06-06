package utils

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// get current absolute path of the running file (go build or go run)
func GetCurrentAbPath() string {
	dir := GetCurrentAbPathByExecutable()
	tmpDir, _ := filepath.EvalSymlinks(os.TempDir())
	if strings.Contains(dir, tmpDir) {
		return GetCurrentAbPathByCaller()
	}
	return dir
}

// get current absolute path of the running file (go build)
func GetCurrentAbPathByExecutable() string {
	exePath, _ := os.Executable()
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

// get current absolute path of the running file (go run)
func GetCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}
