package model

type ADRStatus string

const (
	StatusUnknown    ADRStatus = "Unknown"
	StatusProposed   ADRStatus = "Proposed"
	StatusAccepted   ADRStatus = "Accepted"
	StatusDeprecated ADRStatus = "Deprecated"
	StatusSuperseded ADRStatus = "Superseded"
)

type Adr struct {
	Number int
	Title  string
	Date   string
	Status ADRStatus
}

type AdrConfig struct {
	BaseDir    string `json:"base_directory"`
	CurrentAdr int    `json:"current_id"`
}
