package main

import (
	"os"
	"path/filepath"
)

const DEFAULT_CONFIG_FOLDER_NAME = ".adr"

const DEFAULT_CONFIG_FILE_NAME = "config.json"
const DEFAULT_TEMPLATE_FILE_NAME = "template.md"

var configFolderPath = filepath.Join(getWorkDir(), DEFAULT_CONFIG_FOLDER_NAME)
var configFilePath = filepath.Join(configFolderPath, DEFAULT_CONFIG_FILE_NAME)
var templateFilePath = filepath.Join(configFolderPath, DEFAULT_TEMPLATE_FILE_NAME)

func getWorkDir() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(ex)

}
