package model

type ADRStatus string

const (
	StatusUnknown    ADRStatus = "Unknown"
	StatusProposed   ADRStatus = "Proposed"
	StatusAccepted   ADRStatus = "Accepted"
	StatusDeprecated ADRStatus = "Deprecated"
	StatusSuperseded ADRStatus = "Superseded"
)

type ADR struct {
	Number int
	Title  string
	Date   string
	Status ADRStatus
}
