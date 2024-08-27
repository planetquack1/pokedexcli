package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func getAndPrintPage(cfg *Config) {

	url := createURL(cfg)

	// get bytes from cache, if in cache
	body, ok := cfg.cache.Get(url)

	// if not in cache, HTTP
	if !ok {
		// GET
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		body, err = io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}
		if err != nil {
			log.Fatal(err)
		}
	}

	// Unmarshal based on endpoint

	switch cfg.commandType {
	case "map":

		fmt.Printf("Page number: %d\n", cfg.page+1)

		page := MapPage{}

		err := json.Unmarshal(body, &page)
		if err != nil {
			fmt.Println(err)
		}

		for _, result := range page.Results {
			fmt.Println(result.Name)
		}

	case "explore":
		page := ExplorePage{}

		err := json.Unmarshal(body, &page)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Found Pokemon:")
		for _, encounter := range page.PokemonEncounters {
			fmt.Printf("- %s\n", encounter.Pokemon.Name)
		}
	default:
		fmt.Println("invalid command type")

	}

}
