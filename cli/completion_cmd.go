package cli

import (
	"fmt"
	"strings"
)

const topLevelCompletionWords = "init new list show edit last drop-last completion --help -h --version -v"
const dropLastCompletionWords = "--yes --help -h"
const completionShellWords = "bash zsh fish --help -h"

// CLI Command
type CompletionCmd struct {
	Bash CompletionBashCmd `cmd:"" help:"Generate Bash completion script"`
	Zsh  CompletionZshCmd  `cmd:"" help:"Generate Zsh completion script"`
	Fish CompletionFishCmd `cmd:"" help:"Generate Fish completion script"`
}

type CompletionBashCmd struct{}
type CompletionZshCmd struct{}
type CompletionFishCmd struct{}

func (r *CompletionBashCmd) Run() error {
	fmt.Print(bashCompletionScript())
	return nil
}

func (r *CompletionZshCmd) Run() error {
	fmt.Print(zshCompletionScript())
	return nil
}

func (r *CompletionFishCmd) Run() error {
	fmt.Print(fishCompletionScript())
	return nil
}

func bashCompletionScript() string {
	return fmt.Sprintf(`# bash completion for adr
_adr_completion() {
  local cur prev words cword
  cur="${COMP_WORDS[COMP_CWORD]}"
  prev=""
  if [[ ${COMP_CWORD} -gt 0 ]]; then
    prev="${COMP_WORDS[COMP_CWORD-1]}"
  fi

  if [[ ${COMP_CWORD} -eq 1 ]]; then
    COMPREPLY=( $(compgen -W "%s" -- "${cur}") )
    return 0
  fi

  case "${COMP_WORDS[1]}" in
    drop-last)
      COMPREPLY=( $(compgen -W "%s" -- "${cur}") )
      return 0
      ;;
    completion)
      COMPREPLY=( $(compgen -W "%s" -- "${cur}") )
      return 0
      ;;
  esac
}

complete -F _adr_completion adr
`, topLevelCompletionWords, dropLastCompletionWords, completionShellWords)
}

func zshCompletionScript() string {
	return fmt.Sprintf(`#compdef adr

_adr() {
  local -a commands
  commands=(%s)

  if (( CURRENT == 2 )); then
    _describe 'command' commands
    return
  fi

  case "${words[2]}" in
    drop-last)
      _arguments '%s'
      ;;
    completion)
      _arguments '%s'
      ;;
  esac
}

_adr "$@"
`, zshWords(topLevelCompletionWords), zshArguments(dropLastCompletionWords), zshArguments(completionShellWords))
}

func fishCompletionScript() string {
	return fmt.Sprintf(`# fish completion for adr
complete -c adr -f

complete -c adr -n '__fish_use_subcommand' -a 'init' -d 'Setup ADR directory in the current project'
complete -c adr -n '__fish_use_subcommand' -a 'new' -d 'Creates new ADR using template'
complete -c adr -n '__fish_use_subcommand' -a 'list' -d 'Lists all existing ADRs'
complete -c adr -n '__fish_use_subcommand' -a 'show' -d 'Shows one ADR by number or slug'
complete -c adr -n '__fish_use_subcommand' -a 'edit' -d 'Opens one ADR in an editor'
complete -c adr -n '__fish_use_subcommand' -a 'last' -d 'Shows the newest ADR'
complete -c adr -n '__fish_use_subcommand' -a 'drop-last' -d 'Deletes the newest ADR if it is not in a final state'
complete -c adr -n '__fish_use_subcommand' -a 'completion' -d 'Generate shell completion scripts'
complete -c adr -n '__fish_use_subcommand' -s h -l help -d 'Show context-sensitive help'
complete -c adr -n '__fish_use_subcommand' -s v -l version -d 'Print version information and quit'

complete -c adr -n '__fish_seen_subcommand_from drop-last' -l yes -d 'Delete without confirmation'
complete -c adr -n '__fish_seen_subcommand_from completion; and not __fish_seen_subcommand_from bash zsh fish' -a 'bash zsh fish'
`)
}

func zshWords(words string) string {
	return "'" + replaceSpaces(words, "' '") + "'"
}

func zshArguments(words string) string {
	parts := splitWords(words)
	result := ""
	for i, part := range parts {
		if i > 0 {
			result += " "
		}
		if len(part) > 1 && part[:2] == "--" {
			result += fmt.Sprintf("'%s[%s]'", part, part)
			continue
		}
		result += fmt.Sprintf("'%s'", part)
	}
	return result
}

func splitWords(words string) []string {
	var out []string
	for _, word := range strings.Fields(words) {
		out = append(out, word)
	}
	return out
}

func replaceSpaces(s, replacement string) string {
	return strings.Join(strings.Fields(s), replacement)
}
