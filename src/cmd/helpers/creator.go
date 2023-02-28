package helpers

import (
	"encoding/json"
	"faceitAI/src/models"
	"faceitAI/src/models/dto"
	"fmt"
	"sync"
	"time"
)

var Creator iCreator = creator{}

type iCreator interface {
	CreateMatch(data []byte) models.Match
	CreateMatchListDaysAgo(data []byte, days int) []string
	CreateStats(player *models.Player, wg *sync.WaitGroup)
	CreateStatsToCSV(player *models.Player) string
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

func (c creator) CreateStats(player *models.Player, wg *sync.WaitGroup) {
	defer wg.Done()

	matches := Finder.FindMatchesStats7Days(player.Id)

	Operations.CreateAverageMatch(player, matches)
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

func (c creator) CreateStatsToCSV(player *models.Player) string {
	matches := Finder.FindMatchesStats7Days(player.Id)

	Operations.CreateAverageMatch(player, matches)

	res := fmt.Sprintf("%d, %d, %d, %d, %d, %d, %d, %d, %f, %f, %f, %f",
		player.Stats.Kills,
		player.AverageTeamStats.Kills,
		player.Stats.Assists,
		player.AverageTeamStats.Assists,
		player.Stats.Deaths,
		player.AverageTeamStats.Deaths,
		player.Stats.HSPercent,
		player.AverageTeamStats.HSPercent,
		player.Stats.KillPerRound,
		player.Stats.KillDeathRating,
		player.AverageTeamStats.KillPerRound,
		player.AverageTeamStats.KillDeathRating)

	return res
}
