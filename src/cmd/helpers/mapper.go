package helpers

import (
	"faceitAI/src/models"
	"faceitAI/src/models/dto"
	"strconv"
	"time"
)

var Mapper iMapper = mapper{}

type iMapper interface {
	MatchDto_To_Match(dto.MatchDto) models.Match
	TeamDto_To_Team(dto dto.TeamsDto) models.Team
	PlayersDto_To_Players(dto []dto.PlayersDto) []models.Player
	PlayerDto_To_Player(dto dto.PlayersDto) models.Player

	LiveMatchDto_To_LiveMatch(dto dto.LiveMatchDto) models.LiveMatch
	Faction_To_Team(dto dto.Faction) models.Team
	RosterPlayer_To_Player(dto dto.Roster) models.Player
}

type mapper struct {
	iMapper
}

func (m mapper) MatchDto_To_Match(dto dto.MatchDto) models.Match {
	var match models.Match

	match.Map = dto[0].I1
	match.Date = time.Unix(dto[0].Date/1000, 0)

	match.TeamA = m.TeamDto_To_Team(dto[0].Teams[0])
	match.TeamB = m.TeamDto_To_Team(dto[0].Teams[1])

	return match
}

func (m mapper) TeamDto_To_Team(dto dto.TeamsDto) models.Team {
	var team models.Team

	team.Players = m.PlayersDto_To_Players(dto.Players)
	team.Win = dto.I17

	return team
}

func (m mapper) PlayersDto_To_Players(dto []dto.PlayersDto) []models.Player {
	var players []models.Player

	for _, player := range dto {
		players = append(players, m.PlayerDto_To_Player(player))
	}

	return players
}

func (m mapper) PlayerDto_To_Player(dto dto.PlayersDto) models.Player {
	var player models.Player

	player.Id = dto.PlayerID
	player.Nickname = dto.Nickname

	player.Stats = m.PlayerDto_To_Stats(dto)

	return player
}

func (m mapper) PlayerDto_To_Stats(dto dto.PlayersDto) models.MatchStats {
	var value float64
	var stats models.MatchStats

	value, _ = strconv.ParseFloat(dto.C3, 32)
	stats.KillPerRound = float32(value)

	value, _ = strconv.ParseFloat(dto.C2, 32)
	stats.KillDeathRating = float32(value)

	value, _ = strconv.ParseFloat(dto.I7, 32)
	stats.MVP = float32(value)
	return stats
}

// ---------LIVE MATCH----------
func (m mapper) LiveMatchDto_To_LiveMatch(dto dto.LiveMatchDto) models.LiveMatch {
	var liveMatch models.LiveMatch

	liveMatch.TeamA = m.Faction_To_Team(dto.Payload.Teams.Faction1)
	liveMatch.TeamB = m.Faction_To_Team(dto.Payload.Teams.Faction2)

	return liveMatch
}

func (m mapper) Faction_To_Team(dto dto.Faction) models.Team {
	var team models.Team

	team.EloAverage = strconv.Itoa(dto.Stats.Rating)
	team.TeamName = dto.Name
	for _, player := range dto.Roster {
		team.Players = append(team.Players, m.RosterPlayer_To_Player(player))
	}

	return team
}

func (m mapper) RosterPlayer_To_Player(dto dto.Roster) models.Player {
	var player models.Player

	player.Avatar = dto.Avatar
	player.Id = dto.ID
	player.Level = dto.GameSkillLevel
	player.Nickname = dto.Nickname

	return player
}
