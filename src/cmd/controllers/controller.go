package controllers

import (
	"faceitAI/src/cmd/helpers"
	"net/http"
	"sync"
)

func AnalyzeMatch(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Authorize", "Bearer cd3d93e5-815d-48ec-b6d9-3b77166a3399")

	matchId := r.URL.Query().Get("q")
	match := helpers.Finder.FindMatch(matchId)

	var wg sync.WaitGroup
	wg.Add(1)

	go helpers.Creator.CreateStats(&match.TeamA.Players[0], &wg)
	//...
	//...

	wg.Wait()
}
