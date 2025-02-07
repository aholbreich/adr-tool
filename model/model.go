package model

type AdrStatus string

// TODO
const (
	PROPOSED   AdrStatus = "Proposed"
	ACCEPTED   AdrStatus = "Accepted"
	DEPRECATED AdrStatus = "Deprecated"
	SUPERSEDED AdrStatus = "Superseded"
)

type Adr struct {
	Number int
	Title  string
	Date   string
	Status AdrStatus
}

type AdrConfig struct {
	BaseDir    string `json:"base_directory"`
	CurrentAdr int    `json:"current_id"`
}
