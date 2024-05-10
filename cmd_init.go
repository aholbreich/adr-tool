package main

import (
	"encoding/json"
	"os"

	"github.com/fatih/color"
)

// CLI Command
type InitCmd struct {
}

// Command Handler
func (r *InitCmd) Run(ctx *Globals) error {
	color.Green("Initializing ADR configuration at " + configFolderPath)
	initBaseDir(configFolderPath)
	initConfig(configFolderPath)
	initTemplate()
	return nil
}

func initBaseDir(baseDir string) {
	if _, err := os.Stat(baseDir); os.IsNotExist(err) {
		os.Mkdir(baseDir, 0744)
	} else {
		color.Yellow("Directory" + baseDir + " already exists.")
	}
}

func initConfig(baseDir string) {
	config := AdrConfig{baseDir, 0}
	bytes, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		panic(err)
	}
	color.White("Writngi ")
	os.WriteFile(configFilePath, bytes, 0644)
}

func initTemplate() {
	body := []byte(TEMPLATE1)

	os.WriteFile(templateFilePath, body, 0644)
}
