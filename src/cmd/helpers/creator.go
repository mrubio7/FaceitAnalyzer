package helpers

import (
	"encoding/json"
	"faceitAI/src/models"
	"faceitAI/src/models/dto"
	"sync"
	"time"
)

var Creator iCreator = creator{}

type iCreator interface {
	CreateMatch(data []byte) models.Match
	CreateLiveMatch(data []byte) models.LiveMatch
	CreateMatchListDaysAgo(data []byte, days int) []string
	CreateStats(player *models.Player, wg *sync.WaitGroup)
}

type creator struct {
	iCreator
}

func (c creator) CreateMatch(data []byte) models.Match {
	var matchDto dto.MatchDto
	var match models.Match
	json.Unmarshal(data, &matchDto)

	match = Mapper.MatchDto_To_Match(matchDto)

	return match
}

var P int = 0

func (c creator) CreateStats(player *models.Player, wg *sync.WaitGroup) {
	defer wg.Done()

	matches := Finder.FindMatchesStats7Days(player.Id)

	Operations.CreateAverageMatch(player, matches)

	// fmt.Printf(" P%d ", P)
	// P++
	// if P == 10 {
	// 	P = 0
	// 	fmt.Printf("\n")
	// }
}

func (c creator) CreateMatchListDaysAgo(data []byte, days int) []string {
	var listMatchesDto dto.ListMatchesDto
	var listMatches []string
	json.Unmarshal(data, &listMatchesDto)

	for i, dto := range listMatchesDto {

		date := time.Unix(dto.Date/1000, 0)

		if time.Now().Before(date.AddDate(0, 0, -days)) {
			if i == 0 {
				listMatches = append(listMatches, dto.MatchID)
			}
			break
		}

		listMatches = append(listMatches, dto.MatchID)
	}

	return listMatches
}

func (c creator) CreateLiveMatch(data []byte) models.LiveMatch {
	var matchDto dto.LiveMatchDto
	var match models.LiveMatch
	json.Unmarshal(data, &matchDto)

	match = Mapper.LiveMatchDto_To_LiveMatch(matchDto)

	return match

}
