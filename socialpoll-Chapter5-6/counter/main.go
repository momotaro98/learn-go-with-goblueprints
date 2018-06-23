package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gopkg.in/mgo.v2"
)

var fatalErr error

func fatal(e error) {
	fmt.Println(e)
	flag.PrintDefaults()
	fatalErr = e
}

func main() {
	defer func() { // This defer will be called at the end of main
		if fatalErr != nil {
			os.Exit(1)
		}
	}()

	log.Println("Connecting to DB...")
	db, err := mgo.Dial("localhost")
	if err != nil {
		fatal(err)
		return
	}
	defer func() {
		log.Println("Disconnecting DB...")
		db.Close()
	}()
	pollData := db.DB("ballots").C("polls")
}
