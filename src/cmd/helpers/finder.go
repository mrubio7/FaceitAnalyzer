package helpers

import (
	apiurls_constants "faceitAI/src/constants/apiUrls_constants"
	"faceitAI/src/models"
	"io"
	"log"
	"net/http"
	"strings"
)

var Finder iFinder = finder{}

type iFinder interface {
	FindMatchesStats7Days(playerId string) []models.Match
	FindMatch(matchId string) models.Match
}

type finder struct {
	iFinder
}

func (f finder) FindMatchesStats7Days(playerId string) []models.Match {
	var matches []models.Match

	response, err := http.Get(apiurls_constants.GetMatchList1 + playerId + apiurls_constants.GetMatchList2 + "20")
	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	matchesId := Creator.CreateMatchListDaysAgo(data, 3)

	for _, id := range matchesId {
		match := f.FindMatch(id)
		matches = append(matches, match)
	}

	return matches
}

func (f finder) FindMatch(matchId string) models.Match {
	response, err := http.Get(apiurls_constants.GetMatchStats + matchId)
	if err != nil || response.Status == "404" {
		log.Fatal(err)
	}
	if strings.Contains(response.Status, "429") {
		log.Fatalln("ERROR - Demasiadas peticiones :(")
		log.Panic()
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	match := Creator.CreateMatch(data)

	return match
}
