package helpers

import (
	apiurls_constants "faceitAI/src/constants/apiUrls_constants"
	"faceitAI/src/models"
	"fmt"
	"os/exec"
)

var Finder iFinder = finder{}

type iFinder interface {
	FindMatchesStats7Days(playerId string) []models.Match
	FindMatch(matchId string) models.Match
	FindLiveMatch(matchId string) models.LiveMatch
}

type finder struct {
	iFinder
}

func (f finder) FindMatchesStats7Days(playerId string) []models.Match {
	var matches []models.Match

	// response, err := http.Get(apiurls_constants.GetMatchList1 + playerId + apiurls_constants.GetMatchList2 + "20")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// data, err := io.ReadAll(response.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	curl := exec.Command("curl", "--header", "Host:api.faceit.com", "--header", "Authorization:'Bearer 3a839030-b132-4733-8815-da9a478e835a'", apiurls_constants.GetMatchList1+playerId+apiurls_constants.GetMatchList2+"20")
	out, err := curl.Output()
	if err != nil {
		fmt.Println("erorr", err)
	}

	matchesId := Creator.CreateMatchListDaysAgo(out, 3)

	for _, id := range matchesId {
		match := f.FindMatch(id)
		matches = append(matches, match)
	}

	return matches
}

func (f finder) FindMatch(matchId string) models.Match {
	// response, err := http.Get(apiurls_constants.GetMatchStats + matchId)
	// if err != nil || response.Status == "404" {
	// 	log.Fatal(err)
	// }
	// if strings.Contains(response.Status, "429") {
	// 	log.Fatalln("ERROR - Demasiadas peticiones :(")
	// 	log.Panic()
	// }

	// data, err := io.ReadAll(response.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	curl := exec.Command("curl", "--header", "Host:api.faceit.com", "--header", "Authorization:'Bearer 3a839030-b132-4733-8815-da9a478e835a'", apiurls_constants.GetMatchStats+matchId)
	out, err := curl.Output()
	if err != nil {
		fmt.Println("erorr", err)
	}

	match := Creator.CreateMatch(out)

	return match
}

func (f finder) FindLiveMatch(matchId string) models.LiveMatch {
	// req, _ := http.NewRequest("GET", apiurls_constants.GetMatchLive+matchId, nil)

	// req.Host = "api.faceit.com"
	// req.Header.Add("Authorization", "Bearer 3a839030-b132-4733-8815-da9a478e835a")

	// response, err := http.DefaultClient.Do(req)

	curl := exec.Command("curl", "--header", "Host:api.faceit.com", "--header", "Authorization:'Bearer 3a839030-b132-4733-8815-da9a478e835a'", apiurls_constants.GetMatchLive+matchId)
	out, err := curl.Output()
	if err != nil {
		fmt.Println("erorr", err)
	}

	// if err != nil || response.Status == "404" {
	// 	log.Fatal(err)
	// }
	// if strings.Contains(response.Status, "429") {
	// 	log.Fatalln("ERROR - Demasiadas peticiones :(")
	// 	log.Panic()
	// }

	//data, err := io.ReadAll(response.Body)

	match := Creator.CreateLiveMatch(out)

	return match
}
