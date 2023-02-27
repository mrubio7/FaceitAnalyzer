package controllers

import (
	"faceitAI/src/cmd/helpers"
	apiurls_constants "faceitAI/src/constants/apiUrls_constants"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func FindMatchStats(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Authorize", "Bearer cd3d93e5-815d-48ec-b6d9-3b77166a3399")

	matchId := r.URL.Query().Get("q")
	response, err := http.Get(apiurls_constants.GetMatchStats + matchId)
	if err != nil || response.Status == "404" {
		log.Fatal("Partida no encontrada")
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Partida no encontrada")
	}

	match := helpers.BuildMatch(matchId, data)
	io.WriteString(w, match.Map)
}
