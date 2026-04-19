package cli

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/aholbreich/adr-tool/internal/model"
)

func TestDropLastDeletesNonFinalADR(t *testing.T) {
	oldListADRs := listADRs
	oldLastADRPath := lastADRPath
	oldRemoveADRFile := removeADRFile
	oldConfirmDropLast := confirmDropLast
	t.Cleanup(func() {
		listADRs = oldListADRs
		lastADRPath = oldLastADRPath
		removeADRFile = oldRemoveADRFile
		confirmDropLast = oldConfirmDropLast
	})

	tempDir := t.TempDir()
	target := filepath.Join(tempDir, "003-draft.md")
	if err := os.WriteFile(target, []byte("draft"), 0644); err != nil {
		t.Fatalf("write ADR: %v", err)
	}

	listADRs = func() ([]model.ADR, error) {
		return []model.ADR{{Number: 3, Title: "003-draft.md", Status: model.StatusProposed}}, nil
	}
	lastADRPath = func() (string, error) { return target, nil }
	confirmDropLast = func(prompt string) bool { return true }

	if err := (&DropLastCmd{}).Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	if _, err := os.Stat(target); !os.IsNotExist(err) {
		t.Fatalf("expected ADR to be deleted, stat err = %v", err)
	}
}

func TestDropLastRefusesFinalADR(t *testing.T) {
	tests := []struct {
		name   string
		status model.ADRStatus
	}{
		{name: "accepted", status: model.StatusAccepted},
		{name: "rejected", status: model.StatusRejected},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldListADRs := listADRs
			oldLastADRPath := lastADRPath
			oldRemoveADRFile := removeADRFile
			oldConfirmDropLast := confirmDropLast
			t.Cleanup(func() {
				listADRs = oldListADRs
				lastADRPath = oldLastADRPath
				removeADRFile = oldRemoveADRFile
				confirmDropLast = oldConfirmDropLast
			})

			listADRs = func() ([]model.ADR, error) {
				return []model.ADR{{Number: 3, Title: "003-final.md", Status: tt.status}}, nil
			}
			lastADRPath = func() (string, error) {
				t.Fatal("lastADRPath() should not be called for final ADRs")
				return "", nil
			}
			removeADRFile = func(name string) error {
				t.Fatal("removeADRFile() should not be called for final ADRs")
				return nil
			}

			err := (&DropLastCmd{}).Run()
			if err == nil {
				t.Fatal("expected Run() to fail")
			}
		})
	}
}

func TestDropLastRespectsConfirmation(t *testing.T) {
	oldListADRs := listADRs
	oldLastADRPath := lastADRPath
	oldRemoveADRFile := removeADRFile
	oldConfirmDropLast := confirmDropLast
	t.Cleanup(func() {
		listADRs = oldListADRs
		lastADRPath = oldLastADRPath
		removeADRFile = oldRemoveADRFile
		confirmDropLast = oldConfirmDropLast
	})

	tempDir := t.TempDir()
	target := filepath.Join(tempDir, "003-draft.md")
	if err := os.WriteFile(target, []byte("draft"), 0644); err != nil {
		t.Fatalf("write ADR: %v", err)
	}

	listADRs = func() ([]model.ADR, error) {
		return []model.ADR{{Number: 3, Title: "003-draft.md", Status: model.StatusProposed}}, nil
	}
	lastADRPath = func() (string, error) { return target, nil }
	confirmDropLast = func(prompt string) bool { return false }
	removeADRFile = func(name string) error {
		t.Fatal("removeADRFile() should not be called when deletion is aborted")
		return nil
	}

	if err := (&DropLastCmd{}).Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	if _, err := os.Stat(target); err != nil {
		t.Fatalf("expected ADR to remain after aborted deletion: %v", err)
	}
}

func TestDropLastHandlesEmptyList(t *testing.T) {
	oldListADRs := listADRs
	t.Cleanup(func() {
		listADRs = oldListADRs
	})

	listADRs = func() ([]model.ADR, error) {
		return []model.ADR{}, nil
	}

	err := (&DropLastCmd{}).Run()
	if err == nil {
		t.Fatal("expected Run() to fail for an empty ADR list")
	}
}

func TestDropLastPropagatesDeleteError(t *testing.T) {
	oldListADRs := listADRs
	oldLastADRPath := lastADRPath
	oldRemoveADRFile := removeADRFile
	oldConfirmDropLast := confirmDropLast
	t.Cleanup(func() {
		listADRs = oldListADRs
		lastADRPath = oldLastADRPath
		removeADRFile = oldRemoveADRFile
		confirmDropLast = oldConfirmDropLast
	})

	listADRs = func() ([]model.ADR, error) {
		return []model.ADR{{Number: 3, Title: "003-draft.md", Status: model.StatusProposed}}, nil
	}
	lastADRPath = func() (string, error) { return "/tmp/003-draft.md", nil }
	confirmDropLast = func(prompt string) bool { return true }
	removeADRFile = func(name string) error { return errors.New("permission denied") }

	err := (&DropLastCmd{}).Run()
	if err == nil {
		t.Fatal("expected Run() to fail")
	}
}
