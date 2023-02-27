package models

import (
	"time"
)

type Match struct {
	Date  time.Time `json:"date"`
	Map   string    `json:"map"`
	TeamA Team      `json:"teamA"`
	TeamB Team      `json:"teamB"`
}
