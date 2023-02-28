package helpers

import (
	"faceitAI/src/models"
)

var Operations iOperations = operations{}

type iOperations interface {
	CreateAverageMatch(player *models.Player, match []models.Match)
	CalculateMyAverage(myStats *models.MatchStats, matches []models.Match, nickname string)
	CalculateMatchAverage(average *models.MatchStats, matches []models.Match)
}

type operations struct {
	iOperations
}

func (o operations) CreateAverageMatch(player *models.Player, matches []models.Match) {
	o.CalculateMyAverage(&player.Stats, matches, player.Nickname)
	o.CalculateMatchAverage(&player.AverageTeamStats, matches)

}

func (o operations) CalculateMyAverage(myStats *models.MatchStats, matches []models.Match, nickname string) {

	for _, match := range matches {
		for _, player := range match.TeamA.Players {
			if player.Nickname == nickname {
				sumStats(player.Stats, myStats)
				break
			}
		}
		for _, player := range match.TeamA.Players {
			if player.Nickname == nickname {
				sumStats(player.Stats, myStats)
				break
			}
		}
	}

	count := len(matches)
	averageStats(myStats, count)
}

func (o operations) CalculateMatchAverage(average *models.MatchStats, matches []models.Match) {

	for _, match := range matches {
		for _, player := range match.TeamA.Players {
			sumStats(player.Stats, average)
		}
		for _, player := range match.TeamB.Players {
			sumStats(player.Stats, average)
		}
	}
	averageStats(average, 10)

	count := len(matches)
	averageStats(average, count)
}

func sumStats(player models.MatchStats, stats *models.MatchStats) {
	stats.Assists += player.Assists
	stats.Deaths += player.Deaths
	stats.HSPercent += player.HSPercent
	stats.KillDeathRating += player.KillDeathRating
	stats.KillPerRound += player.KillPerRound
	stats.Kills += player.Kills
}

func averageStats(stats *models.MatchStats, num int) {
	stats.Assists /= num
	stats.Deaths /= num
	stats.HSPercent /= num
	stats.KillDeathRating /= float32(num)
	stats.KillPerRound /= float32(num)
	stats.Kills /= num
}
