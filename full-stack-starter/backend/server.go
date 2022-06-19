package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Result struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/click", clickMe)
	log.Println("Starting a web server on port 8080")
	http.ListenAndServe(":8080", nil)
}

func clickMe(w http.ResponseWriter, r *http.Request) {
	input, _ := strconv.Atoi(r.URL.Query().Get("c"))

	var result Result
	message := "The answer is " + strconv.Itoa(input)
	result.Message = message

	output, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println(result)
		return
	}

	// Allows CORS and writes output to the connection as JSON
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
