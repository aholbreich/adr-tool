package adr

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
	"unicode"

	"github.com/aholbreich/adr-tool/internal/config"
	"github.com/aholbreich/adr-tool/internal/model"
	"github.com/aholbreich/adr-tool/internal/template"
)

type AdrManager struct {
}

// NewManager initializes a new ADR Manager
func NewAdrManager() *AdrManager {
	return &AdrManager{}
}

// Create a new ADR file with the given name
func (m *AdrManager) CreateNewAdr(currentConfig model.AdrConfig, adrName string) error {

	adr := model.Adr{
		Title:  adrName,
		Date:   time.Now().Format("2006-01-02 15:04"), // ISO 8601 format
		Number: currentConfig.CurrentAdr,
		Status: model.PROPOSED,
	}

	// Use TemplateManager to load template

	tpl, err := template.NewTplManager().LoadTemplate()
	if err != nil {
		return fmt.Errorf("failed to load template: %w", err)
	}

	// Sanitize and build the filename
	adrFileName := fmt.Sprintf("%03d-%s.md", adr.Number, strings.Join(strings.Split(strings.TrimSpace(adr.Title), " "), "-"))
	adrFullPath := filepath.Join(currentConfig.BaseDir, adrFileName)

	// Create and write to the ADR file
	f, err := os.Create(adrFullPath)
	if err != nil {
		return fmt.Errorf("failed to create ADR file: %w", err)
	}
	defer f.Close()

	// Render template into the new ADR file
	if err := tpl.Execute(f, adr); err != nil {
		return fmt.Errorf("failed to render template: %w", err)
	}

	return nil
}

// List ADR files in reverse order
func (m *AdrManager) GetADRList() ([]string, error) {

	entries, err := os.ReadDir(config.PathResolverInst().ConfigFolderPath())
	if err != nil {
		return nil, err
	}

	var adrs []string
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if len(name) > 0 && unicode.IsDigit(rune(name[0])) { // Starts with digit? Must be an ADR file
			adrs = append(adrs, name)
		}
	}

	sort.Sort(sort.Reverse(sort.StringSlice(adrs))) // Reverse order
	return adrs, nil
}
