# ADR Tool Go

Opinionated version of ADR tool written in go. Initially inspired by the [adr-tools](https://github.com/npryce/adr-tools).


# Installation 

## RPM (Fedora, RedHat)

First Add RPM repository
```bash
# Docu: https://aholbreich.github.io/rpm-repo/#installation-fedora-centos-redhat
echo '[Holbreich]
name=Holbreich Repository
baseurl=https://aholbreich.github.io/rpm-repo/
enabled=1
gpgcheck=0' | sudo tee /etc/yum.repos.d/holbreich.repo

```
install rpm with `yum` or `dnf`

```bash

sudo dnf install adr-tool

```
Find support in [rpm-repo](https://github.com/aholbreich/rpm-repo) project home in case of any issues.

## Local build

The straightforward way to compile it on your own.

```bash
git clone https://github.com/aholbreich/adr-tool.git
cd adr-tool
# compile and put to $(HOME)/bin
make install

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
Shows you list of your ADRs with corresponding status

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

* [x] Add build pipeline
* [x] Add ADR Status Info in listing
* [x] Multi platform binaries
* [ ] Too long being not in final status warning
* [ ] Add Status transition?
* [ ] Color codes?
* [ ] Release notes (See https://github.com/git-chglog/git-chglog)


## For developers

```bash
# VBuild and try local
make build

make test

make clean
```