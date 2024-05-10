package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

// CLI Command
type NewCmd struct {
	Name string `arg required help:"ADR Name"`
}

// Command Handler
func (r *NewCmd) Run() error {
	fmt.Println(" Creating new ADR" + r.Name)
	currentConfig := getConfig()
	currentConfig.CurrentAdr++
	updateConfig(currentConfig)
	newAdr(currentConfig, r.Name)
	return nil

}

func getConfig() AdrConfig {
	var currentConfig AdrConfig

	bytes, err := os.ReadFile(configFilePath)
	if err != nil {
		color.Red("No ADR configuration is found!")
		color.HiGreen("Start by initializing ADR configuration, check 'adr init --help' for more help")
		os.Exit(1)
	}

	json.Unmarshal(bytes, &currentConfig)
	return currentConfig
}

func updateConfig(config AdrConfig) {
	bytes, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		panic(err)
	}
	os.WriteFile(configFilePath, bytes, 0644)
}

func newAdr(config AdrConfig, adrName string) {
	adr := Adr{
		Title:  strings.Join([]string{adrName}, " "),
		Date:   time.Now().Format("02-01-2006 15:04:05"),
		Number: config.CurrentAdr,
		Status: PROPOSED,
	}
	template, err := template.ParseFiles(templateFilePath)
	if err != nil {
		panic(err)
	}
	adrFileName := strconv.Itoa(adr.Number) + "-" + strings.Join(strings.Split(strings.Trim(adr.Title, "\n \t"), " "), "-") + ".md"
	adrFullPath := filepath.Join(config.BaseDir, adrFileName)
	f, err := os.Create(adrFullPath)
	if err != nil {
		panic(err)
	}
	template.Execute(f, adr)
	f.Close()
	color.Green("ADR number " + strconv.Itoa(adr.Number) + " was successfully written to : " + adrFullPath)
}
