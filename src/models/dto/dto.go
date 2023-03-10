package dto

type MatchDto []struct {
	ID            string     `json:"_id,omitempty"`
	CreatedAt     int64      `json:"created_at,omitempty"`
	UpdatedAt     int64      `json:"updated_at,omitempty"`
	BestOf        string     `json:"bestOf,omitempty"`
	CompetitionID string     `json:"competitionId,omitempty"`
	Date          int64      `json:"date,omitempty"`
	Game          string     `json:"game,omitempty"`
	GameMode      string     `json:"gameMode,omitempty"`
	I0            string     `json:"i0,omitempty"`
	I1            string     `json:"i1,omitempty"`
	I12           string     `json:"i12,omitempty"`
	I18           string     `json:"i18,omitempty"`
	I2            string     `json:"i2,omitempty"`
	MatchID       string     `json:"matchId,omitempty"`
	MatchRound    string     `json:"matchRound,omitempty"`
	Played        string     `json:"played,omitempty"`
	Teams         []TeamsDto `json:"teams,omitempty"`
}
type PlayersDto struct {
	I9       string `json:"i9,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	I10      string `json:"i10,omitempty"`
	I13      string `json:"i13,omitempty"`
	I15      string `json:"i15,omitempty"`
	I6       string `json:"i6,omitempty"`
	I14      string `json:"i14,omitempty"`
	I7       string `json:"i7,omitempty"`
	I16      string `json:"i16,omitempty"`
	I8       string `json:"i8,omitempty"`
	PlayerID string `json:"playerId,omitempty"`
	C3       string `json:"c3,omitempty"`
	C2       string `json:"c2,omitempty"`
	C4       string `json:"c4,omitempty"`
}
type TeamsDto struct {
	I19     string       `json:"i19,omitempty"`
	Players []PlayersDto `json:"players,omitempty"`
	TeamID  string       `json:"teamId,omitempty"`
	I3      string       `json:"i3,omitempty"`
	I4      string       `json:"i4,omitempty"`
	I5      string       `json:"i5,omitempty"`
	Premade bool         `json:"premade,omitempty"`
	I17     string       `json:"i17,omitempty"`
	C9      string       `json:"c9,omitempty"`
	C5      string       `json:"c5,omitempty"`
}

type ListMatchesDto []struct {
	ID struct {
		MatchID  string `json:"matchId"`
		PlayerID string `json:"playerId"`
	} `json:"_id"`
	CreatedAt     int64  `json:"created_at"`
	UpdatedAt     int64  `json:"updated_at"`
	I9            string `json:"i9"`
	Nickname      string `json:"nickname"`
	I10           string `json:"i10"`
	I13           string `json:"i13"`
	I15           string `json:"i15"`
	I6            string `json:"i6"`
	I14           string `json:"i14"`
	I7            string `json:"i7"`
	I16           string `json:"i16"`
	I8            string `json:"i8"`
	PlayerID      string `json:"playerId"`
	C3            string `json:"c3"`
	C2            string `json:"c2"`
	C4            string `json:"c4"`
	C1            string `json:"c1"`
	I19           string `json:"i19"`
	TeamID        string `json:"teamId"`
	I3            string `json:"i3"`
	I4            string `json:"i4"`
	I5            string `json:"i5"`
	Premade       bool   `json:"premade"`
	C5            string `json:"c5"`
	BestOf        string `json:"bestOf"`
	CompetitionID string `json:"competitionId"`
	Date          int64  `json:"date"`
	Game          string `json:"game"`
	GameMode      string `json:"gameMode"`
	I0            string `json:"i0"`
	I1            string `json:"i1"`
	I12           string `json:"i12"`
	I18           string `json:"i18"`
	I2            string `json:"i2"`
	MatchID       string `json:"matchId"`
	MatchRound    string `json:"matchRound"`
	Played        string `json:"played"`
	Status        string `json:"status"`
	Elo           string `json:"elo"`
}
