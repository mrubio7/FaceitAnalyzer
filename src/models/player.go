package models

type Player struct {
	Id         string     `json:"playerId"`
	Nickname   string     `json:"nickname"`
	Stats      MatchStats `json:"stats"`
	TeamStats  MatchStats `json:"teamStats"`
	EnemyStats MatchStats `json:"enemyStats"`
}

type MatchStats struct {
	KillPerRound    float32 `json:"c3"`
	KillDeathRating float32 `json:"c2"`
	MVP             float32 `json:"c4"`
}
