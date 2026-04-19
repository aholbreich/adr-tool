package cli

import (
	"strings"
	"testing"
)

func TestBashCompletionScriptIncludesCommands(t *testing.T) {
	script := bashCompletionScript()

	for _, want := range []string{"drop-last", "completion", "--version", "bash zsh fish"} {
		if !strings.Contains(script, want) {
			t.Fatalf("bash completion script missing %q", want)
		}
	}
}

func TestZshCompletionScriptIncludesCommands(t *testing.T) {
	script := zshCompletionScript()

	for _, want := range []string{"#compdef adr", "drop-last", "completion", "--version"} {
		if !strings.Contains(script, want) {
			t.Fatalf("zsh completion script missing %q", want)
		}
	}
}

func TestFishCompletionScriptIncludesCommands(t *testing.T) {
	script := fishCompletionScript()

	for _, want := range []string{"complete -c adr", "drop-last", "completion", "-l version"} {
		if !strings.Contains(script, want) {
			t.Fatalf("fish completion script missing %q", want)
		}
	}
}
