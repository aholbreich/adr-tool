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
)

var adrFileNamePattern = regexp.MustCompile(`^(\d+)-.+\.md$`)
var invalidSlugCharsPattern = regexp.MustCompile(`[^a-z0-9-]+`)
var repeatedHyphensPattern = regexp.MustCompile(`-+`)
var statusLinePattern = regexp.MustCompile(`(?im)^\s*(?:[*-]\s*)?Status:\s*([A-Za-z]+)\b`)

type ADRManager struct{}

// NewADRManager initializes a new ADR manager.
func NewADRManager() *ADRManager {
	return &ADRManager{}
}

// CreateNewADR creates a new ADR file with the given name and returns its full path.
func (m *ADRManager) CreateNewADR(baseDir string, number int, adrName string) (string, error) {
	adr := model.ADR{
		Title:  adrName,
		Date:   time.Now().Format("2006-01-02 15:04"),
		Number: number,
		Status: model.StatusProposed,
	}

	tpl, err := template.NewTplManager().LoadTemplate()
	if err != nil {
		return "", fmt.Errorf("failed to load template: %w", err)
	}

	adrFileName := buildADRFileName(adr.Number, adr.Title)
	adrFullPath := filepath.Join(baseDir, adrFileName)

	if err := os.MkdirAll(baseDir, 0755); err != nil {
		return "", fmt.Errorf("failed to ensure ADR directory exists: %w", err)
	}

	f, err := os.OpenFile(adrFullPath, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
	if err != nil {
		return "", fmt.Errorf("failed to create ADR file: %w", err)
	}
	defer f.Close()

	if err := tpl.Execute(f, adr); err != nil {
		return "", fmt.Errorf("failed to render template: %w", err)
	}

	return adrFullPath, nil
}

// NextADRNumber returns the next ADR number based on the existing files in baseDir.
func (m *ADRManager) NextADRNumber(baseDir string) (int, error) {
	adrs, err := m.listADRsInDir(baseDir)
	if err != nil {
		if os.IsNotExist(err) {
			return 1, nil
		}
		return 0, err
	}

	maxNumber := 0
	for _, adr := range adrs {
		if adr.Number > maxNumber {
			maxNumber = adr.Number
		}
	}

	return maxNumber + 1, nil
}

// ListADRs returns all ADR files in reverse numeric order.
func (m *ADRManager) ListADRs() ([]model.ADR, error) {
	configDir := config.PathResolverInst().ConfigFolderPath()
	return m.listADRsInDir(configDir)
}

func (m *ADRManager) listADRsInDir(configDir string) ([]model.ADR, error) {
	entries, err := os.ReadDir(configDir)
	if err != nil {
		if os.IsNotExist(err) {
			return []model.ADR{}, nil
		}
		return nil, err
	}

	var adrs []model.ADR
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

		adrs = append(adrs, model.ADR{
			Number: num,
			Title:  name,
			Status: status,
		})
	}

	SortADRListReverse(adrs)
	return adrs, nil
}

func extractStatus(filePath string) (model.ADRStatus, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return model.StatusUnknown, err
	}

	matches := statusLinePattern.FindStringSubmatch(string(content))
	if len(matches) > 1 {
		return parseADRStatus(matches[1]), nil
	}

	return model.StatusUnknown, nil
}

func parseADRStatus(value string) model.ADRStatus {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case strings.ToLower(string(model.StatusProposed)):
		return model.StatusProposed
	case strings.ToLower(string(model.StatusAccepted)):
		return model.StatusAccepted
	case strings.ToLower(string(model.StatusDeprecated)):
		return model.StatusDeprecated
	case strings.ToLower(string(model.StatusSuperseded)):
		return model.StatusSuperseded
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
