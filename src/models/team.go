package models

type Team struct {
	Players []Player `json:"players"`
	Predict Predict  `json:"predict"`
}
