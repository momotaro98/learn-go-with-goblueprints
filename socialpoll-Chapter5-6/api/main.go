package main

import (
	"flag"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
)

func main() {
	var (
		addr  = flag.String("addr", ":8080", "endpoint address")
		mongo = flag.String("mongo", "localhost", "mongodb address")
	)
	log.Println("Dialing mongo", *mongo)
	db, err := mgo.Dial(*mongo)
	if err != nil {
		log.Fatalln("failed to connect to mongo:", err)
	}
	defer db.Close()
	log.Println(db)
	s := &Server{
		db: db,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/polls/", withCORS(withAPIKey(s.handlePolls)))
	log.Println(mux)
	log.Println("Starting web server on", *addr)
	http.ListenAndServe(":8080", mux)
	log.Println("Stopping....")
}

// Server is the API server.
type Server struct {
	db *mgo.Session
}

func withCORS(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Expose-Headers", "Location")
		fn(w, r)
	}
}

// withAPIKey wraps HandlerFunc
func withAPIKey(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !isValidAPIKey(r.URL.Query().Get("key")) {
			respondErr(w, r, http.StatusUnauthorized, "invalid API key")
			return
		}
		fn(w, r)
	}
}

func isValidAPIKey(key string) bool {
	return key == "abc123"
}
