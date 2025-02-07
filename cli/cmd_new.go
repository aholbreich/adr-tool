package cli

import (
	"encoding/json"
	"html/template"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"adr-tool/model"

	"github.com/fatih/color"
)

// CLI Command
type NewCmd struct {
	AdrName []string `arg:"" required:"" help:"ADR Name"`
}

// Command Handler
func (r *NewCmd) Run() error {
	adrName := strings.Join(r.AdrName, " ")
	currentConfig := getConfig()
	currentConfig.CurrentAdr++
	updateConfig(currentConfig)
	newAdr(currentConfig, adrName)
	return nil

}

func getConfig() model.AdrConfig {
	var currentConfig model.AdrConfig

	bytes, err := os.ReadFile(configFilePath)
	if err != nil {
		color.Red("No ADR configuration is found!")
		color.HiGreen("Start by initializing ADR configuration, check 'adr init --help' for more help")
		os.Exit(1)
	}

	json.Unmarshal(bytes, &currentConfig)
	return currentConfig
}

func updateConfig(config model.AdrConfig) {
	bytes, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		panic(err)
	}
	os.WriteFile(configFilePath, bytes, 0644)
}

func newAdr(config model.AdrConfig, adrName string) {
	adr := model.Adr{
		Title:  strings.Join([]string{adrName}, " "),
		Date:   time.Now().Format("01-02-2006 15:04"),
		Number: config.CurrentAdr,
		Status: model.PROPOSED,
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
	color.Green("New ADR " + strconv.Itoa(adr.Number) + " was successfully written to : " + adrFullPath)
}
