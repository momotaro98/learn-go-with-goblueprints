package main

import (
	"encoding/json"
	"net/http"

	"github.com/momotaro98/learn-go-with-goblueprints/meander-Chapter7/meander"
)

func main() {
	meander.APIKey = "YOUR API KEY HERE" // TODO: set API key by env var
	http.HandleFunc("/journeys", func(w http.ResponseWriter, r *http.Request) {
		respond(w, r, meander.Journeys)
	})
	http.ListenAndServe(":8080", http.DefaultServeMux)
}

func respond(w http.ResponseWriter, r *http.Request, data []interface{}) error {
	publicData := make([]interface{}, len(data))
	for i, d := range data {
		publicData[i] = meander.Public(d)
	}
	return json.NewEncoder(w).Encode(publicData)
}
