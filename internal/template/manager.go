package template

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"os"
)

//go:embed templates/default.md
var embeddedTemplates embed.FS

type TemplateManager struct {
}

// NewManager initializes a new Template Manager
func NewManager() *TemplateManager {
	return &TemplateManager{}
}

// LoadTemplate loads the embedded template
func (m *TemplateManager) LoadTemplate() (*template.Template, error) {
	// Access the embedded file
	templateData, err := fs.ReadFile(embeddedTemplates, "templates/default.md")
	if err != nil {
		return nil, err
	}

	// Parse the template
	return template.New("adr").Parse(string(templateData))
}

// SaveTemplate writes the rendered template to the specified file
func (m *TemplateManager) SaveProcessed(filePath string, data interface{}) error {
	tmpl, err := m.LoadTemplate()
	if err != nil {
		return fmt.Errorf("failed to load template: %w", err)
	}

	f, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to write template: %w", err)
	}
	defer f.Close()
	return tmpl.Execute(f, data)
}
