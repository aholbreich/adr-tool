package cli

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
)

const DEFAULT_CONFIG_FOLDER_NAME = ".adr"

const DEFAULT_CONFIG_FILE_NAME = "config.json"
const DEFAULT_TEMPLATE_FILE_NAME = "template.md"

var configFolderPath = filepath.Join(getWorkDir(), DEFAULT_CONFIG_FOLDER_NAME)
var configFilePath = filepath.Join(configFolderPath, DEFAULT_CONFIG_FILE_NAME)
var templateFilePath = filepath.Join(configFolderPath, DEFAULT_TEMPLATE_FILE_NAME)

// Filename is the __filename equivalent
func Filename() (string, error) {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return "", errors.New("unable to get the current filename")
	}
	return filename, nil
}

func getWorkDir() string {

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return dir
}
