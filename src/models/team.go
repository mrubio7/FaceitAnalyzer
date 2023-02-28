package models

type Team struct {
	Players    []Player `json:"players"`
	WinPercent float32  `json:"winPercent"`
}
