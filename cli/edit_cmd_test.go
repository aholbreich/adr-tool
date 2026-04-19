package cli

import (
	"errors"
	"testing"
)

func TestResolveEditorUsesVISUALBeforeEDITOR(t *testing.T) {
	oldLookupEnv := lookupEnv
	oldLookPath := lookPath
	t.Cleanup(func() {
		lookupEnv = oldLookupEnv
		lookPath = oldLookPath
	})

	lookupEnv = func(key string) (string, bool) {
		switch key {
		case "VISUAL":
			return "visual-editor", true
		case "EDITOR":
			return "plain-editor", true
		default:
			return "", false
		}
	}
	lookPath = func(file string) (string, error) {
		if file == "visual-editor" {
			return "/usr/bin/visual-editor", nil
		}
		return "", errors.New("not found")
	}

	got, err := resolveEditor()
	if err != nil {
		t.Fatalf("resolveEditor() unexpected error: %v", err)
	}

	if got != "visual-editor" {
		t.Fatalf("got %q, want %q", got, "visual-editor")
	}
}

func TestResolveEditorFallsBackToDefaultEditors(t *testing.T) {
	oldLookupEnv := lookupEnv
	oldLookPath := lookPath
	oldCurrentGOOS := currentGOOS
	t.Cleanup(func() {
		lookupEnv = oldLookupEnv
		lookPath = oldLookPath
		currentGOOS = oldCurrentGOOS
	})

	lookupEnv = func(key string) (string, bool) {
		return "", false
	}
	currentGOOS = "linux"
	lookPath = func(file string) (string, error) {
		if file == "vi" {
			return "/usr/bin/vi", nil
		}
		return "", errors.New("not found")
	}

	got, err := resolveEditor()
	if err != nil {
		t.Fatalf("resolveEditor() unexpected error: %v", err)
	}

	if got != "vi" {
		t.Fatalf("got %q, want %q", got, "vi")
	}
}

func TestResolveEditorReturnsHelpfulErrorWhenMissing(t *testing.T) {
	oldLookupEnv := lookupEnv
	oldLookPath := lookPath
	oldCurrentGOOS := currentGOOS
	t.Cleanup(func() {
		lookupEnv = oldLookupEnv
		lookPath = oldLookPath
		currentGOOS = oldCurrentGOOS
	})

	lookupEnv = func(key string) (string, bool) {
		return "", false
	}
	currentGOOS = "linux"
	lookPath = func(file string) (string, error) {
		return "", errors.New("not found")
	}

	_, err := resolveEditor()
	if err == nil {
		t.Fatal("expected resolveEditor() to fail")
	}
}
