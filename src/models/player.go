package models

type Player struct {
	Id           string     `json:"playerId"`
	Nickname     string     `json:"nickname"`
	MyStats      MatchStats `json:"stats"`
	AverageStats MatchStats `json:"averageStats"`
}

type MatchStats struct {
	Win             bool    `json:"i17"`
	NumRounds       int     `json:"c5"`
	Kills           int     `json:"i6"`
	Deaths          int     `json:"i8"`
	Assists         int     `json:"i7"`
	KillPerRound    float32 `json:"c3"`
	KillDeathRating float32 `json:"c2"`
	HSPercent       int     `json:"c1"`
	MVP             int     `json:"i9"`
	TripleKills     int     `json:"i14"`
	QuadroKills     int     `json:"i15"`
	PentaKills      int     `json:"i16"`
	EnemyAverageElo int     `json:"enemyAverageElo"`
	Elo             int     `json:"elo"`
	GamesToday      int     `json:"gamesToday"`
}
