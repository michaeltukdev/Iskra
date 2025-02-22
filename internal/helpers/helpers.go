package helpers

import (
	"log"
	"path/filepath"
	"runtime"
)

func GetProjectRoot() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Unable to get the current file path")
	}

	projectRoot := filepath.Dir(filepath.Dir(filepath.Dir(filename)))

	return projectRoot
}
