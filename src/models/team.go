package models

type Team struct {
	TeamName   string   `json:"teamName"`
	EloAverage string   `json:"averageElo"`
	Players    []Player `json:"players"`
	Win        string   `json:"win"`
}
