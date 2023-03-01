package helpers

import (
	"faceitAI/src/models"
)

var Operations iOperations = operations{}

type iOperations interface {
	CreateAverageMatch(player *models.Player, match []models.Match)
	CalculateMyAverage(myStats *models.MatchStats, matches []models.Match, nickname string)
	CalculateMatchAverage(average *models.MatchStats, matches []models.Match)
	GetStats(players []models.Player) models.MatchStats
}

type operations struct {
	iOperations
}

func (o operations) CreateAverageMatch(player *models.Player, matches []models.Match) {
	o.CalculateMyAverage(&player.Stats, matches, player.Nickname)
	o.CalculateEnemyAverage(&player.EnemyStats, matches, player.Nickname)
	o.CalculateTeamAverage(&player.TeamStats, matches, player.Nickname)
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

func (o operations) CalculateEnemyAverage(average *models.MatchStats, matches []models.Match, nick string) {
	var stats []models.MatchStats

	for _, match := range matches {
		for _, player := range match.TeamA.Players {
			var stat models.MatchStats

			if player.Nickname != nick {
				stat = GetStats(match.TeamA.Players)
			} else {
				stat = GetStats(match.TeamB.Players)
			}

			stats = append(stats, stat)
		}
	}
}

func (o operations) CalculateTeamAverage(average *models.MatchStats, matches []models.Match, nick string) {
	var stats []models.MatchStats

	for _, match := range matches {
		for _, player := range match.TeamA.Players {
			var stat models.MatchStats

			if player.Nickname == nick {
				stat = GetStats(match.TeamA.Players)
			} else {
				stat = GetStats(match.TeamB.Players)
			}

			stats = append(stats, stat)
		}
	}
}

func GetStats(players []models.Player) models.MatchStats {
	var stat models.MatchStats

	for _, player := range players {
		sumStats(player.Stats, &stat)
	}

	averageStats(&stat, 5)

	return stat
}

func sumStats(player models.MatchStats, stats *models.MatchStats) {
	stats.KillDeathRating += player.KillDeathRating
	stats.KillPerRound += player.KillPerRound
	stats.MVP += player.MVP
}

func averageStats(stats *models.MatchStats, num int) {
	stats.KillDeathRating /= float32(num)
	stats.KillPerRound /= float32(num)
	stats.MVP /= float32(num)
}
