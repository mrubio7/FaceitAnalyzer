package controllers

import (
	"encoding/csv"
	"faceitAI/src/cmd/helpers"
	"faceitAI/src/models"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func AnalyzeMatch(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Authorize", "Bearer cd3d93e5-815d-48ec-b6d9-3b77166a3399")

	matchId := r.URL.Query().Get("q")
	//match := helpers.Finder.FindMatch(matchId)
	liveMatch := helpers.Finder.FindLiveMatch(matchId)

	var wg sync.WaitGroup

	wg.Add(10)

	go helpers.Creator.CreateStats(&liveMatch.TeamA.Players[0], &wg)
	go helpers.Creator.CreateStats(&liveMatch.TeamA.Players[1], &wg)
	go helpers.Creator.CreateStats(&liveMatch.TeamA.Players[2], &wg)
	go helpers.Creator.CreateStats(&liveMatch.TeamA.Players[3], &wg)
	go helpers.Creator.CreateStats(&liveMatch.TeamA.Players[4], &wg)

	go helpers.Creator.CreateStats(&liveMatch.TeamB.Players[0], &wg)
	go helpers.Creator.CreateStats(&liveMatch.TeamB.Players[1], &wg)
	go helpers.Creator.CreateStats(&liveMatch.TeamB.Players[2], &wg)
	go helpers.Creator.CreateStats(&liveMatch.TeamB.Players[3], &wg)
	go helpers.Creator.CreateStats(&liveMatch.TeamB.Players[4], &wg)

	wg.Wait()

	var averageTeamA models.Player
	var averageTeamB models.Player

	for _, player := range liveMatch.TeamA.Players {
		sumStats(player.Stats, &averageTeamA.Stats)
		sumStats(player.EnemyStats, &averageTeamA.EnemyStats)
		sumStats(player.TeamStats, &averageTeamA.TeamStats)
	}
	for _, player := range liveMatch.TeamB.Players {
		sumStats(player.Stats, &averageTeamB.Stats)
		sumStats(player.EnemyStats, &averageTeamB.EnemyStats)
		sumStats(player.TeamStats, &averageTeamB.TeamStats)
	}

	averageStats(&averageTeamA.EnemyStats, 5)
	averageStats(&averageTeamA.TeamStats, 5)
	averageStats(&averageTeamA.Stats, 5)

	averageStats(&averageTeamB.EnemyStats, 5)
	averageStats(&averageTeamB.TeamStats, 5)
	averageStats(&averageTeamB.Stats, 5)

	var result models.Results

	result.Data = append(result.Data, float64(averageTeamA.EnemyStats.KillDeathRating))
	result.Data = append(result.Data, float64(averageTeamA.EnemyStats.KillPerRound))
	result.Data = append(result.Data, float64(averageTeamA.EnemyStats.MVP))

	result.Data = append(result.Data, float64(averageTeamA.TeamStats.KillDeathRating))
	result.Data = append(result.Data, float64(averageTeamA.TeamStats.KillPerRound))
	result.Data = append(result.Data, float64(averageTeamA.TeamStats.MVP))

	result.Data = append(result.Data, float64(averageTeamA.Stats.KillDeathRating))
	result.Data = append(result.Data, float64(averageTeamA.Stats.KillPerRound))
	result.Data = append(result.Data, float64(averageTeamA.Stats.MVP))

	result.Data = append(result.Data, float64(averageTeamB.EnemyStats.KillDeathRating))
	result.Data = append(result.Data, float64(averageTeamB.EnemyStats.KillPerRound))
	result.Data = append(result.Data, float64(averageTeamB.EnemyStats.MVP))

	result.Data = append(result.Data, float64(averageTeamB.TeamStats.KillDeathRating))
	result.Data = append(result.Data, float64(averageTeamB.TeamStats.KillPerRound))
	result.Data = append(result.Data, float64(averageTeamB.TeamStats.MVP))

	result.Data = append(result.Data, float64(averageTeamB.Stats.KillDeathRating))
	result.Data = append(result.Data, float64(averageTeamB.Stats.KillPerRound))
	result.Data = append(result.Data, float64(averageTeamB.Stats.MVP))

	prediction := helpers.Analyze.Predict(result)

	fmt.Fprintf(w, "%f", prediction)
}

func CreateCSV(w http.ResponseWriter, r *http.Request) {
	log.Println("OK   - Initializing CreateCSV")
	r.Header.Set("Authorize", "Bearer cd3d93e5-815d-48ec-b6d9-3b77166a3399")
	//r.Header.Set("Authorize", "Bearer 5650365f-4b6e-41d0-90d3-f9a3bcb5dccb")

	matches := []string{
		"1-57093027-40d8-4005-8d0d-2a0db30bdbbf",
		"1-c91c5b52-93d3-4912-a4fc-400cd7a089cf",
		"1-a7e211d2-babf-4e7f-96fc-36c1cd314863",
		"1-b5b76c5b-b26c-4809-9c8a-04908677547c",
		"1-42960fe5-50a8-4188-b847-d9d3109852dd",
		"1-0c7c4a43-955c-4584-b5d8-865c9dff5a58",
		"1-3b8d2e1c-ecef-41c5-9666-e0ddf639420c",
		"1-0c1f7923-02e8-4881-ba8b-c0d618b3a3d8",
		"1-863c1135-4eae-4f42-8318-fa11ac4d786e",
		"1-bfa78781-6ef8-4a0a-adcd-165db0a9a6b8",
		"1-4d3e9bad-41ec-4504-9c37-e9b73b8920fa",
		"1-00038c2c-a69d-4462-b310-e9094cee4842",
		"1-3e245a8b-f5aa-4f60-bea7-2ace19c02a44",
		"1-a8a7ecb6-4f82-4ea1-9aa4-e1a0554de4f4",
		"1-a5f5f71e-d442-45a9-a4c7-1874ac68e2b1",
		"1-6a065d89-3395-4e1f-9cb5-e997de96b6a4",
		"1-31259c50-096e-4f80-85f3-55770da32a8a",
		"1-fed8677c-7404-4137-a56b-8535d91a232c",
		"1-87044324-78e7-4721-b4cc-876cd0da6a05",
		"1-a2cf426e-18e6-44d1-966b-e07efd909b14",
		"1-6a065d89-3395-4e1f-9cb5-e997de96b6a4",
		"1-5ea2a916-9dc7-4ea6-8470-0521cf01fe02",
		"1-93b528a1-4e46-4e0c-acd7-f4c0dfe8f683",
		"1-8565ca3b-01dd-4b8a-a8fa-ad725755924f",
		"1-6c9b9b08-69a5-417f-b17d-43cc58047eaa",
		"1-eb0e1f8f-74ec-4b27-a265-65bd155c72e7",
		"1-2377f4f0-f83b-4926-8ebf-d5280422e6a1",
		"1-2884a66f-3730-4f49-bf68-0f79f2b2fe7b",
		"1-d16e0176-d236-47b0-ad17-b72c9508e0ae",
		"1-6e4b9550-c8c8-4a07-877b-948612a6e254",
		"1-7ceba21a-7692-4d69-999f-e033a669ab9d",
		"1-52cd0dd7-1d26-4546-b016-1a968f81ea59",
		"1-1b25af1c-bcae-49e6-bbd3-c0cbce3b2339",
		"1-fa4f6573-507d-4c96-b902-648633cfeb0c",
		"1-1d6b115e-0c42-484c-9cfc-174e8687bb8d",
		"1-8b7cd53d-9c5a-4eb0-b676-3d148afa7b05",
		"1-f35e9c04-8b7e-463a-9d8a-41aab3f6d10f",
		"1-66d52409-ab17-42f0-a64a-5624febe9712",
		"1-33ed3ed6-0dd1-4f79-9112-37e31f29157d",
		"1-85691ea8-9cc8-448d-b9be-06e71eb54757",
		"1-fd561f57-e104-4f18-bd74-2929d296723b",
		"1-df539030-2925-4224-b2d9-dc028b43ee30",
		"1-4f88a09e-53f7-4bb4-b178-3393fadef100",
		"1-25428bf3-1b63-4465-b830-debfaa7fad0d",
		"1-78e3322d-cd92-4441-843d-440f9be2bae8",
		"1-7a2428d1-fb70-4efe-9d0c-8b98cb7bd62a",
		"1-338bc6fa-5a92-435e-8883-11c204fe4f1c",
		"1-0bdf344f-7d76-4696-a384-048c38873266",
		"1-55a75b5d-16a5-4b3e-bdc1-812652767ef4",
		"1-4c02ac2e-a9eb-4596-968c-503c86b1182b",
		"1-38db4206-ee81-423a-8ab8-e213c8f2c04d",
		"1-569ff570-193a-419a-b840-bbc896df57f5",
		"1-df255e52-48b8-4959-bc3d-bb190f8c5933",
		"1-2a271cb0-12bd-43bb-ac01-65ff8a673e39",
		"1-28aeeaf9-d6c2-4142-94b7-903b5ceed61e",
		"1-95435ef4-182b-4f1b-ae28-dd9bdb87a51e",
		"1-1ef00bcf-7dda-4d0f-b926-647296ecd276",
		"1-b7a8cdbc-bc62-45b1-bef3-fe00e914ca86",
		"1-e6cdb57d-6944-48d9-9c1f-93a9bea44f39",
		"1-1dba4c60-6e3b-45e0-abab-b31e39fbcda6",
		"1-2c8f697a-4aab-44ca-8f0c-f7eecf777571",
		"1-e24c8c02-a5ea-4e13-9c9b-eb5bd81e8586",
		"1-04ccd9b3-6e82-4dc7-b5ba-13e149fe8b5f",
		"1-b3d9fa8f-6aaf-4c43-904f-fde73bc6c91a",
	}

	f, err := os.Create("data/data.csv")
	if err != nil {
		log.Fatal(err)
	}

	headers := fmt.Sprintf("%s, %s, %s\n",
		createHeader("TeamA"),
		createHeader("TeamB"),
		"TeamA_Wins")

	f.WriteString(headers)

	for _, matchId := range matches {
		match := helpers.Finder.FindMatch(matchId)

		log.Printf("WORK - Analyzing match (%s)\n", matchId)

		var wg sync.WaitGroup
		wg.Add(10)
		go helpers.Creator.CreateStats(&match.TeamA.Players[0], &wg)
		go helpers.Creator.CreateStats(&match.TeamA.Players[1], &wg)
		go helpers.Creator.CreateStats(&match.TeamA.Players[2], &wg)
		go helpers.Creator.CreateStats(&match.TeamA.Players[3], &wg)
		go helpers.Creator.CreateStats(&match.TeamA.Players[4], &wg)

		go helpers.Creator.CreateStats(&match.TeamB.Players[0], &wg)
		go helpers.Creator.CreateStats(&match.TeamB.Players[1], &wg)
		go helpers.Creator.CreateStats(&match.TeamB.Players[2], &wg)
		go helpers.Creator.CreateStats(&match.TeamB.Players[3], &wg)
		go helpers.Creator.CreateStats(&match.TeamB.Players[4], &wg)
		wg.Wait()

		var averageTeamA models.Player
		var averageTeamB models.Player

		for _, player := range match.TeamA.Players {
			sumStats(player.Stats, &averageTeamA.Stats)
			sumStats(player.EnemyStats, &averageTeamA.EnemyStats)
			sumStats(player.TeamStats, &averageTeamA.TeamStats)
		}
		for _, player := range match.TeamB.Players {
			sumStats(player.Stats, &averageTeamB.Stats)
			sumStats(player.EnemyStats, &averageTeamB.EnemyStats)
			sumStats(player.TeamStats, &averageTeamB.TeamStats)
		}

		averageStats(&averageTeamA.EnemyStats, 5)
		averageStats(&averageTeamA.TeamStats, 5)
		averageStats(&averageTeamA.Stats, 5)
		fmt.Printf("Averaged TeamA  ")

		averageStats(&averageTeamB.EnemyStats, 5)
		averageStats(&averageTeamB.TeamStats, 5)
		averageStats(&averageTeamB.Stats, 5)
		fmt.Printf("Averaged TeamB\n")

		str := fmt.Sprintf("%s, %s, %s\n",
			createPlayerStat(averageTeamA),
			createPlayerStat(averageTeamB),
			match.TeamA.Win)

		f.WriteString(str)

		log.Printf("OK   - Analyzed match (%s)\n", matchId)

		time.Sleep(10 * time.Second)
	}
}

func Training(w http.ResponseWriter, r *http.Request) {
	fd, _ := os.Open("data/dataTraining.csv")
	f := csv.NewReader(fd)

	records, _ := f.ReadAll()

	var results []models.Results

	for i, match := range records {
		if i == 64 {
			i = 64
		}
		var res models.Results

		if i == 0 {
			continue
		}

		for _, data := range match {
			data = strings.TrimSpace(data)

			d, _ := strconv.ParseFloat(data, 64)
			res.Data = append(res.Data, d)
		}

		label := res.Data[len(res.Data)-1]
		res.Label = int(label)

		res.Data = res.Data[:len(res.Data)-1] //quitamos el ultimo

		results = append(results, res)
	}

	helpers.Analyze.Training(results)
}

func createHeader(h string) string {
	return fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s, %s",
		h+"_KD",
		h+"_KR",
		h+"_MVP",
		h+"_T_KD",
		h+"_T_KR",
		h+"_T_MVP",
		h+"_E_KD",
		h+"_E_KR",
		h+"_E_MVP")
}

func createPlayerStat(pl models.Player) string {
	return fmt.Sprintf("%f, %f, %f, %f, %f, %f, %f, %f, %f",
		pl.Stats.KillDeathRating,
		pl.Stats.KillPerRound,
		pl.Stats.MVP,
		pl.TeamStats.KillDeathRating,
		pl.TeamStats.KillPerRound,
		pl.TeamStats.MVP,
		pl.EnemyStats.KillDeathRating,
		pl.EnemyStats.KillPerRound,
		pl.EnemyStats.MVP,
	)
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
