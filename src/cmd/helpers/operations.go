package helpers

import "faceitAI/src/models"

var Operations iOperations = operations{}

type iOperations interface {
	CreateAverageMatch(player *models.Player, match []models.Match)
}

type operations struct {
	iOperations
}

func (o operations) CreateAverageMatch(player *models.Player, match []models.Match) {

}
