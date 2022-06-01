package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/andrei-k/go-rps-web-app/rps"
)

func main() {
	http.HandleFunc("/play", playRound)
	http.HandleFunc("/", homePage)

	log.Println("Starting a web server on port 8080")
	// Start a production ready web server
	http.ListenAndServe(":8080", nil)
}

func playRound(w http.ResponseWriter, r *http.Request) {
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("c"))
	result := rps.PlayRound(playerChoice)

	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println(result)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")
}

func renderTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)
	if err != nil {
		log.Println(err)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}
}
