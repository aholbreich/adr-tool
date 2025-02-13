package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

const (
	DefaultConfigFolderName = ".adr"
	DefaultConfigFileName   = "config.json"
	DefaultTemplateFileName = "default.md"
)

type PathResolver struct {
	BaseDir string
}

// NewPathResolver initializes a new PathResolver with the working directory
func NewPathResolver() *PathResolver {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: Unable to get the current working directory. Please check your environment and try again.")
		os.Exit(1) // Exit with status code 1 indicating failure
	}
	return &PathResolver{BaseDir: dir}
}

func (p *PathResolver) ConfigFolderPath() string {
	return filepath.Join(p.BaseDir, DefaultConfigFolderName)
}

// ConfigFilePath returns the full path to the config file
func (p *PathResolver) ConfigFilePath() string {
	return filepath.Join(p.ConfigFolderPath(), DefaultConfigFileName)
}

// TemplateFilePath returns the full path to the template file
func (p *PathResolver) TemplateFilePath() string {
	return filepath.Join(p.ConfigFolderPath(), DefaultTemplateFileName)
}

func (p *PathResolver) IsFilepathGitRepo() bool {
	if _, err := os.Stat(filepath.Join(".", ".git")); os.IsNotExist(err) {
		return false
	}
	return true
}

// Filename returns the current file's filename
func Filename() (string, error) {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return "", errors.New("unable to get the current filename")
	}
	return filename, nil
}
