# ADR Tool Go

Yet another ADR Tool written in Go.

Inspired by the [adr-tools](https://github.com/npryce/adr-tools) but using the Go instead of Bash.

# Installation 

## RPM (Fedora, RedHat)

First Add RPM repository
```bash
echo '[aHolbreich Repository]
name=aHolbreich Repository
baseurl=https://aholbreich.github.io/rpm-repo/
enabled=1
gpgcheck=0' | sudo tee /etc/yum.repos.d/aHolbreich.repo

```
install rpm with `yum` or `dnf`

```bash

sudo dnf install adr-tool

```

## Local build

```bash
#Build locally
go build -o adr
# Make available to the system
sudo mv adr /usr/local/bin/adr

# Alternatively make it available to your user only
mv adr ~/bin/adr

```

# Usage

## Init configuration

Run
```bash
adr init 
```
before you start working.

## Creating a new ADR

```bash
adr new how to make CLI tools
```
this will create a new numbered ADR in folder `.adr`:
`1-how to make CLI tools.md`.

## Listing existing ADRs

```bash
adr list 
```

## Help and Docu

```bash
# List all commands
adr -h 

#Example detailed help to a particular subcommand
adr new -h 
```

## Composing, Editing and Change Status of your ADR

User your favored Editor, Open the desired ADR file under ./.adr/ folder change anything.

## TODOs

* [.] Add build pipeline
* [ ] Add ADR Status Info in listing (Require file parsing)
* [ ] Add Status transition? 


## For developers

```bash
# VBuild and try local
go build -o adr

./adr
```