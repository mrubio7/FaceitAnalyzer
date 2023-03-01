package models

import (
	"time"
)

type Match struct {
	Date  time.Time `json:"date"`
	Map   string    `json:"i1"`
	TeamA Team      `json:"teamA"`
	TeamB Team      `json:"teamB"`
	Win   bool      `json:"win"`
}
