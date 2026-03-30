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

var adrFileNamePattern = regexp.MustCompile(`^(\d+)-.+\.md$`)
var invalidSlugCharsPattern = regexp.MustCompile(`[^a-z0-9-]+`)
var repeatedHyphensPattern = regexp.MustCompile(`-+`)

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
		Date:   time.Now().Format("2006-01-02 15:04"),
		Number: currentConfig.CurrentAdr,
		Status: model.StatusProposed,
	}

	tpl, err := template.NewTplManager().LoadTemplate()
	if err != nil {
		return fmt.Errorf("failed to load template: %w", err)
	}

	adrFileName := buildADRFileName(adr.Number, adr.Title)
	adrFullPath := filepath.Join(currentConfig.BaseDir, adrFileName)

	f, err := os.Create(adrFullPath)
	if err != nil {
		return fmt.Errorf("failed to create ADR file: %w", err)
	}
	defer f.Close()

	if err := tpl.Execute(f, adr); err != nil {
		return fmt.Errorf("failed to render template: %w", err)
	}

	return nil
}

func (m *AdrManager) GetADRList() ([]model.Adr, error) {
	configDir := config.PathResolverInst().ConfigFolderPath()
	entries, err := os.ReadDir(configDir)
	if err != nil {
		return nil, err
	}

	var adrs []model.Adr
	for _, e := range entries {
		if e.IsDir() {
			continue
		}

		name := e.Name()
		num, err := extractNumberFromString(name)
		if err != nil {
			continue
		}

		adrPath := filepath.Join(configDir, name)
		status, err := extractStatus(adrPath)
		if err != nil {
			status = model.StatusUnknown
		}

		adrs = append(adrs, model.Adr{
			Number: num,
			Title:  name,
			Status: status,
		})
	}

	SortAdrListReverse(adrs)
	return adrs, nil
}

func extractStatus(filePath string) (model.ADRStatus, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return model.StatusUnknown, err
	}

	re := regexp.MustCompile(`(?i)^Status:\s*(\w+)`)
	matches := re.FindStringSubmatch(string(content))
	if len(matches) > 1 {
		caser := cases.Title(language.English)
		return parseADRStatus(caser.String(matches[1])), nil
	}

	return model.StatusUnknown, nil
}

func parseADRStatus(value string) model.ADRStatus {
	switch model.ADRStatus(value) {
	case model.StatusProposed,
		model.StatusAccepted,
		model.StatusDeprecated,
		model.StatusSuperseded:
		return model.ADRStatus(value)
	default:
		return model.StatusUnknown
	}
}

func extractNumberFromString(s string) (int, error) {
	matches := adrFileNamePattern.FindStringSubmatch(s)
	if len(matches) < 2 {
		return 0, fmt.Errorf("no ADR number found in %q", s)
	}

	number, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0, err
	}

	return number, nil
}

func buildADRFileName(number int, title string) string {
	slug := slugifyTitle(title)
	if slug == "" {
		slug = "untitled"
	}

	return fmt.Sprintf("%03d-%s.md", number, slug)
}

func slugifyTitle(title string) string {
	normalized := strings.ToLower(strings.TrimSpace(title))
	normalized = strings.Join(strings.Fields(normalized), "-")
	normalized = invalidSlugCharsPattern.ReplaceAllString(normalized, "-")
	normalized = repeatedHyphensPattern.ReplaceAllString(normalized, "-")
	return strings.Trim(normalized, "-")
}
