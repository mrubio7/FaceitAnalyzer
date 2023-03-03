package main

import (
	"faceitAI/src/cmd/controllers"
	"log"
	"net/http"
)

func main() {
	go api()

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func api() {

	http.HandleFunc("/analyze", controllers.AnalyzeMatch)
	http.HandleFunc("/data", controllers.CreateCSV)
	http.HandleFunc("/train", controllers.Training)
}
