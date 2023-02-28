package models

type Player struct {
	Id               string     `json:"playerId"`
	Nickname         string     `json:"nickname"`
	Stats            MatchStats `json:"stats"`
	AverageTeamStats MatchStats `json:"averageTeamStats"`
}

type MatchStats struct {
	Kills           int     `json:"i6"`
	Deaths          int     `json:"i8"`
	Assists         int     `json:"i7"`
	KillPerRound    float32 `json:"c3"`
	KillDeathRating float32 `json:"c2"`
	HSPercent       int     `json:"c4"`
}
