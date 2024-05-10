# ADR Tool Go

Yet another ADR Tool written in Go.

Inspired by the [adr-tools](https://github.com/npryce/adr-tools) but using the Go instead of Bash.

# Init configuration

Run
```bash
adr init 
```
before you start working.

## Creating a new ADR

```bash
adr new "ADR Name"
```
this will create a new numbered ADR in folder `.adr`:
`1-ADR Name`.

## Listing existing ADRs

```bash
adr list 
```

## Composing, Editing and Change Status of your ADR

User your favored Editor, Open the desired ADR file under ./.adr/ folder change anything.

## TODOs
[ ] Add build pipeline
[ ] Add ADR Status Info in listing (Require file parsing)
[ ] Add Status transition? 