package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"

	"github.com/mtslzr/pokeapi-go"
)

func getAndPrintResponse(cfg *Config) {

	// Unmarshal based on endpoint

	switch cfg.commandType {
	case "map":

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

		l, err := pokeapi.LocationArea(cfg.location)
		if err != nil {
			fmt.Printf("Error retrieving %s\n", cfg.location)
			break
		}

		fmt.Println("Found Pokemon:")
		for _, encounter := range l.PokemonEncounters {
			fmt.Printf("- %s\n", encounter.Pokemon.Name)
		}

	case "catch":

		var chanceOneIn int
		if cfg.baseExperience < 50 {
			chanceOneIn = 1
		} else {
			chanceOneIn = cfg.baseExperience / 50
		}

		fmt.Printf("Chance of catching is 1 in %d\n", chanceOneIn)

		if rand.Int()%chanceOneIn == 0 {
			fmt.Printf("caught %s!\n", cfg.pokemon.Name)
			addPokemon(cfg)
		} else {
			fmt.Printf("%s escaped!\n", cfg.pokemon.Name)
		}

	default:
		fmt.Println("invalid command type")

	}

}
