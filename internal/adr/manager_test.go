package adr

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/aholbreich/adr-tool/internal/model"
)

func TestExtractNumberFromString(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int
		wantErr bool
	}{
		{name: "valid ADR file", input: "001-example.md", want: 1},
		{name: "valid ADR file with larger number", input: "123-something.md", want: 123},
		{name: "missing dash separator", input: "123example.md", wantErr: true},
		{name: "number not at start", input: "foo-123.md", wantErr: true},
		{name: "not markdown", input: "001-example.txt", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractNumberFromString(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error for %q", tt.input)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got != tt.want {
				t.Fatalf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestParseADRStatus(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  model.ADRStatus
	}{
		{name: "known status", input: "Proposed", want: model.StatusProposed},
		{name: "known status lowercase", input: "proposed", want: model.StatusProposed},
		{name: "another known status", input: "Accepted", want: model.StatusAccepted},
		{name: "unknown status", input: "SomethingElse", want: model.StatusUnknown},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseADRStatus(tt.input)
			if got != tt.want {
				t.Fatalf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestExtractStatus(t *testing.T) {
	tests := []struct {
		name    string
		content string
		want    model.ADRStatus
	}{
		{
			name: "plain status line below title",
			content: `# 1. Example

Status: Proposed
`,
			want: model.StatusProposed,
		},
		{
			name: "markdown bullet status line",
			content: `# 1. Example

* Status: Accepted
`,
			want: model.StatusAccepted,
		},
		{
			name: "missing status line",
			content: `# 1. Example

No status here.
`,
			want: model.StatusUnknown,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempDir := t.TempDir()
			filePath := filepath.Join(tempDir, "001-example.md")
			if err := os.WriteFile(filePath, []byte(tt.content), 0644); err != nil {
				t.Fatalf("write test ADR: %v", err)
			}

			got, err := extractStatus(filePath)
			if err != nil {
				t.Fatalf("extractStatus() unexpected error: %v", err)
			}

			if got != tt.want {
				t.Fatalf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestBuildADRFileName(t *testing.T) {
	tests := []struct {
		name   string
		number int
		title  string
		want   string
	}{
		{name: "normal title", number: 1, title: "How to make CLI tools", want: "001-how-to-make-cli-tools.md"},
		{name: "punctuation and spaces", number: 12, title: "  Hello,   World!  ", want: "012-hello-world.md"},
		{name: "empty title falls back", number: 7, title: "   ", want: "007-untitled.md"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildADRFileName(tt.number, tt.title)
			if got != tt.want {
				t.Fatalf("got %q, want %q", got, tt.want)
			}
		})
	}
}
