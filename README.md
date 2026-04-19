# ADR Tool Go

`adr` is a small command-line tool for creating and listing Architecture Decision Records (ADRs) in a project.

For a general introduction to the topic, see:
[Architecture Decision Records: A Tool for Experienced Engineers](https://alexander.holbreich.org/adr_method/)

## Installation

### Binary download

Prebuilt binaries for Linux, macOS, and Windows are available on the
[Releases page](https://github.com/aholbreich/adr-tool/releases).

### RPM (Fedora / Red Hat)

Add the RPM repository:

```bash
# Documentation: https://aholbreich.github.io/rpm-repo/#installation-fedora-centos-redhat
echo '[Holbreich]
name=Holbreich Repository
baseurl=https://aholbreich.github.io/rpm-repo/
enabled=1
gpgcheck=0' | sudo tee /etc/yum.repos.d/holbreich.repo
```

Install the package:

```bash
sudo dnf install adr-tool
```

If you run into issues with the RPM repository, see the
[rpm-repo project](https://github.com/aholbreich/rpm-repo).

### Local build

Build and install the binary into `$(HOME)/bin`:

```bash
git clone https://github.com/aholbreich/adr-tool.git
cd adr-tool
make install
```

If `$(HOME)/bin` is not in your `PATH`, either add it or set a custom install directory:

```bash
make install INSTALL_DIR=/usr/local/bin
```

## Usage

The adr tool usage is more or less self explainable.

```bash

 adr 
Usage: adr <command> [flags]

ADR tool for your project.

    Project details can be found at https://github.com/aholbreich/adr-tool

Flags:
  -h, --help       Show context-sensitive help.
  -v, --version    Print version information and quit

Commands:
  init          Setup ADR directory in the current project
  new           Creates new ADR using template
  list          Lists all existing ADRs
  show          Shows one ADR by number or slug
  edit          Opens one ADR in an editor
  last          Shows the newest ADR
  drop-last     Deletes the newest ADR if it is not in a final state
  commit        Stages and commits ADR changes only
  completion    Generate shell completion scripts

Run "adr <command> --help" for more information on a command.
```

### Initialize ADR directory
Run this once in the root of your project:

```bash
adr init
```

This creates the `.adr/` directory.

If no `.git` directory is found, the tool warns you and asks for confirmation before continuing.

### Create a new ADR

```bash
adr new how to make cli tools
```

This creates a new numbered ADR file inside `.adr/`, for example:

```text
001-how-to-make-cli-tools.md
```

The next ADR number is derived from the existing ADR files in `.adr/`.

### List existing ADRs

```bash
adr list
```

Example output:

```text
Architecture Decision Records:
 - 003-example-of-rejected [Unknown]
 - 002-comsi-comsa [Proposed]
 - 001-better-folder-structure [Accepted]
```

### Show one ADR

```bash
# By number
adr show 1

# By full stem
adr show 001-how-to-make-cli-tools

# By slug
adr show how-to-make-cli-tools
```

This prints the full ADR content to standard output.

### Edit one ADR

```bash
# By number
adr edit 1

# By slug
adr edit how-to-make-cli-tools
```

This opens the ADR in `$VISUAL`, then `$EDITOR`, then a small OS-specific fallback editor list.

### Show the newest ADR

```bash
adr last
```

This prints the full content of the highest-numbered ADR.

### Delete the newest draft ADR

```bash
# Prompt before deleting
adr drop-last

# Delete without prompting
adr drop-last --yes
```

This deletes the highest-numbered ADR only when its status is not final.

### Commit ADR changes

```bash
# Auto-generated commit message
adr commit

# Explicit commit message
adr commit -m "Add authentication ADR"
```

This stages changes under `.adr/` and creates a Git commit containing only ADR changes. If the current directory is not a Git repository, the command fails with a clear error.

### Generate shell completion

```bash
adr completion bash
adr completion zsh
adr completion fish
```

Example installation:

```bash
# Bash
install -d ~/.local/share/bash-completion/completions 
adr completion bash > ~/.local/share/bash-completion/completions/adr

# Zsh
adr completion zsh > "${fpath[1]}/_adr"

# Fish
adr completion fish > ~/.config/fish/completions/adr.fish
```

### Show help

```bash
# Show top-level help
adr -h

# Show help for a subcommand
adr new -h
```

### Show version

```bash
adr --version
# or
adr -v
```

## Editing ADRs

Use your preferred editor to modify ADR files in the `.adr/` directory.

The default template looks like this:

```md
# <number>. <title>

Status: Draft
Status Date: <date>
Driver: <Your Name>
Contributors: ...

## Context

## Decision

### Consequences

## Options considered

### Option 1:

### Option 2:

## Advices
```

The `adr list` command reads the `Status:` line from each ADR file and shows it in the output.

## Development

Useful local targets:

```bash
# Format, tidy, and build the binary
make build

# Run tests
make test

# Build distribution archives
make dists

# Show computed version string
make get-version

# Remove local build artifacts
make clean

# Clean Go caches as well
make cleancache
```

## Project status / ideas

See also: [ROADMAP.md](./ROADMAP.md)

Implemented:
- multi-platform binaries

