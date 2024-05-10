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
	initConfig(configFolderPath)
	initTemplate()
	return nil
}

func initConfig(baseDir string) {
	if _, err := os.Stat(baseDir); os.IsNotExist(err) {
		os.Mkdir(baseDir, 0744)
	} else {
		color.Red("Directory" + baseDir + " already exists. Not overriding.")
		os.Exit(-1)
	}
	config := AdrConfig{baseDir, 0}
	bytes, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		panic(err)
	}
	color.White("Writing new configuration at: " + configFilePath)
	os.WriteFile(configFilePath, bytes, 0644)
}

func initTemplate() {
	body := []byte(TEMPLATE1)

	os.WriteFile(templateFilePath, body, 0644)
}
