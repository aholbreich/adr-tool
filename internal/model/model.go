package model

// TODO
const (
	PROPOSED   string = "Proposed"
	ACCEPTED   string = "Accepted"
	DEPRECATED string = "Deprecated"
	SUPERSEDED string = "Superseded"
)

type Adr struct {
	Number int
	Title  string
	Date   string
	Status string
}

type AdrConfig struct {
	BaseDir    string `json:"base_directory"`
	CurrentAdr int    `json:"current_id"`
}
