package controllers

import (
	"encoding/csv"
	"encoding/json"
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
	//r.Header.Set("Authorize", "Bearer cd3d93e5-815d-48ec-b6d9-3b77166a3399")
	r.Header.Set("Authorize", "Bearer 5650365f-4b6e-41d0-90d3-f9a3bcb5dccb")
	w.Header().Set("Content-Type", "application/json")

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

	liveMatch.Result = helpers.Analyze.Predict(result)
	liveMatch.ReadableResult = createReadableResult(liveMatch.Result)

	// lm, _ := json.Marshal(liveMatch)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(liveMatch)
}

func AnalyzePreviusMatch(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Authorize", "Bearer cd3d93e5-815d-48ec-b6d9-3b77166a3399")

	matchId := r.URL.Query().Get("q")
	//match := helpers.Finder.FindMatch(matchId)
	liveMatch := helpers.Finder.FindMatch(matchId)

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
		"1-7058e650-e862-4a4f-9080-92c9e7f30708",
		"1-38b292f1-a6f0-484e-a3d5-20deee41c23e",
		"1-1044e5df-9032-4674-9489-7051ce7ba180",
		"1-251b2fe8-f2f1-4a85-8c04-115e5a49e0e4",
		"1-63d03389-4246-4095-b38b-29acfebfcd85",
		"1-7b82548a-73a5-409d-95ac-4f73e54f328a",
		"1-73c8d48a-21d9-460c-99ba-a6c87bcca5b3",
		"1-463cd6af-58ea-4204-a853-dc58da554978",
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
	fd, _ := os.Open("data/dataMerged.csv")
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

func createReadableResult(res float64) string {
	results := []string{
		"16-0",
		"16-1",
		"16-2",
		"16-3",
		"16-4",
		"16-5",
		"16-6",
		"16-7",
		"16-8",
		"16-9",
		"16-10",
		"16-11",
		"16-12",
		"16-13",
		"16-14",
		"14-16",
		"13-16",
		"12-16",
		"11-16",
		"10-16",
		"9-16",
		"8-16",
		"7-16",
		"6-16",
		"5-16",
		"4-16",
		"3-16",
		"2-16",
		"1-16",
		"0-16"}

	temp := 1.0 / float64(len(results))

	index := res / temp

	return results[int(index)]
}
