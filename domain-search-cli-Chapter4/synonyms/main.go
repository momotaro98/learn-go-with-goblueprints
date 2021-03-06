package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/momotaro98/learn-go-with-goblueprints/domain-search-cli-Chapter4/thesaurus"
)

func main() {
	apiKey := os.Getenv("BHT_APIKEY")
	if apiKey == "" {
		log.Fatalln("Failed to get ENVIRONMENT VAR, 'BHT_APIKEY'")
		return
	}
	thesaurus := &thesaurus.BigHugh{APIKey: apiKey}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := s.Text()
		syns, err := thesaurus.Synonyms(word)
		if err != nil {
			log.Fatalln("Failed when looking for synonyms for \""+word+"\"", err)
		}
		if len(syns) == 0 {
			log.Fatalln("Couldn't find any synonyms for \"" + word + "\"")
		}
		for _, syn := range syns {
			fmt.Println(syn)
		}
	}
}
