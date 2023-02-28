package helpers

import (
	"faceitAI/src/models"
	"sync"
)

var Operations iOperations = operations{}

type iOperations interface {
	CreateAverageMatch(player *models.Player, match []models.Match)
	CalculateMyAverage(myStats *models.MatchStats, matches []models.Match, wg *sync.WaitGroup, nickname string)
	CalculateMatchAverage(average *models.MatchStats, matches []models.Match, wg *sync.WaitGroup)
}

type operations struct {
	iOperations
}

func (o operations) CreateAverageMatch(player *models.Player, matches []models.Match) {
	var wg2 sync.WaitGroup

	wg2.Add(2)

	go o.CalculateMyAverage(&player.MyStats, matches, &wg2, player.Nickname)
	go o.CalculateMatchAverage(&player.AverageStats, matches, &wg2)

	wg2.Done()
}

func (o operations) CalculateMyAverage(myStats *models.MatchStats, matches []models.Match, wg *sync.WaitGroup, nickname string) {
	defer wg.Done()

	for _, match := range matches {
		for _, player := range match.TeamA.Players {
			if player.Nickname == nickname {
				sumStats(player.MyStats, myStats)
				break
			}
		}
		for _, player := range match.TeamA.Players {
			if player.Nickname == nickname {
				sumStats(player.MyStats, myStats)
				break
			}
		}
	}

	count := len(matches)
	averageStats(myStats, count)
}

func (o operations) CalculateMatchAverage(average *models.MatchStats, matches []models.Match, wg *sync.WaitGroup) {
	defer wg.Done()

	for _, match := range matches {
		for _, player := range match.TeamA.Players {
			sumStats(player.AverageStats, average)
		}
		for _, player := range match.TeamB.Players {
			sumStats(player.AverageStats, average)
		}
		averageStats(average, 10)
	}
	count := len(matches)
	averageStats(average, count)
}

func sumStats(player models.MatchStats, stats *models.MatchStats) {
	stats.Assists += player.Assists
	stats.Deaths += player.Deaths
	stats.Elo += player.Elo
	stats.EnemyAverageElo += player.EnemyAverageElo
	stats.HSPercent += player.HSPercent
	stats.KillDeathRating += player.KillDeathRating
	stats.KillPerRound += player.KillPerRound
	stats.Kills += player.Kills
	stats.MVP += player.MVP
	stats.NumRounds += player.NumRounds
	stats.PentaKills += player.PentaKills
	stats.QuadroKills += player.QuadroKills
	stats.TripleKills += player.TripleKills
}

func averageStats(stats *models.MatchStats, num int) {
	stats.Assists /= num
	stats.Deaths /= num
	stats.Elo /= num
	stats.EnemyAverageElo /= num
	stats.HSPercent /= num
	stats.KillDeathRating /= float32(num)
	stats.KillPerRound /= float32(num)
	stats.Kills /= num
	stats.MVP /= num
	stats.NumRounds /= num
	stats.PentaKills /= num
	stats.QuadroKills /= num
	stats.TripleKills /= num
}
