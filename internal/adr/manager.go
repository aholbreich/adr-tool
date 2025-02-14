package adr

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/aholbreich/adr-tool/internal/config"
	"github.com/aholbreich/adr-tool/internal/model"
	"github.com/aholbreich/adr-tool/internal/template"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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

func (m *AdrManager) GetADRList() ([]model.Adr, error) {
	entries, err := os.ReadDir(config.PathResolverInst().ConfigFolderPath())
	if err != nil {
		return nil, err
	}

	var adrs []model.Adr
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if len(name) > 0 {
			num, err := extractNumberFromString(name)
			if err != nil {
				continue // not having number is not ADR file.
			}

			adrPath := config.PathResolverInst().ConfigFolderPath() + "/" + name
			status, err := extractStatus(adrPath)
			if err != nil {
				status = "Unknown" // Default to "Unknown" if status cannot be extracted
			}

			adr := model.Adr{
				Number: num,
				Title:  name,
				Status: status,
			}
			adrs = append(adrs, adr)
		}
	}

	SortAdrListReverse(adrs)

	return adrs, nil
}

func extractStatus(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// Regex to find the status, e.g., "Status: Accepted"
	re := regexp.MustCompile(`(?i)Status:\s*(\w+)`)
	matches := re.FindStringSubmatch(string(content))
	if len(matches) > 1 {
		// Capitalize first letter using x/text/cases
		caser := cases.Title(language.English)
		return caser.String(matches[1]), nil
	}

	return "Unknown", nil
}

func extractNumberFromString(s string) (int, error) {

	re := regexp.MustCompile(`\d+`) // \d+ matches one or more digits
	matches := re.FindString(s)

	if matches == "" {
		return 0, fmt.Errorf("no number found in the string")
	}

	// Convert the matched string to an integer
	number, err := strconv.Atoi(matches)
	if err != nil {
		return 0, err
	}

	return number, nil
}
