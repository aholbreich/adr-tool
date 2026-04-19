package model

type ADRStatus string

const (
	StatusUnknown    ADRStatus = "Unknown"
	StatusDraft      ADRStatus = "Draft"
	StatusProposed   ADRStatus = "Proposed"
	StatusAccepted   ADRStatus = "Accepted"
	StatusRejected   ADRStatus = "Rejected"
	StatusDeprecated ADRStatus = "Deprecated"
	StatusSuperseded ADRStatus = "Superseded"
)

type ADR struct {
	Number int
	Title  string
	Date   string
	Status ADRStatus
}

func IsFinalStatus(status ADRStatus) bool {
	switch status {
	case StatusAccepted, StatusRejected, StatusDeprecated, StatusSuperseded:
		return true
	default:
		return false
	}
}
