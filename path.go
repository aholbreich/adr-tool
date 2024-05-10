package main

import (
	"errors"
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

// Dirname is the __dirname equivalent
func getWorkDir() string {

	filename, err := Filename()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(filename)
}
