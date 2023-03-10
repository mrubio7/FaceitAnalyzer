package models

import (
	"time"
)

type Match struct {
	Date  time.Time `json:"date"`
	Map   string    `json:"i1"`
	TeamA Team      `json:"teamA"`
	TeamB Team      `json:"teamB"`
}

type LiveMatch struct {
	ReadableResult string
	Result         float64
	TeamA          Team
	TeamB          Team
}
