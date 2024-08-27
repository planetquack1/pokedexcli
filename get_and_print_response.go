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

		// Unmarshal
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
		if cfg.pokemon.BaseExperience < 50 {
			chanceOneIn = 1
		} else {
			chanceOneIn = cfg.pokemon.BaseExperience / 50
		}

		fmt.Printf("Chance of catching is 1 in %d\n", chanceOneIn)

		if rand.Int()%chanceOneIn == 0 {
			fmt.Printf("caught %s!\n", cfg.pokemon.Name)
			addPokemon(cfg)
		} else {
			fmt.Printf("%s escaped!\n", cfg.pokemon.Name)
		}

	case "inspect":

		p, ok := cfg.pokedex[cfg.pokemon.Name]

		if !ok {
			fmt.Printf("you have not yet caught %s\n", cfg.pokemon.Name)
			return
		}

		// Print stats
		fmt.Printf("Name: %s\n", p.Name)
		fmt.Printf("Height: %d\n", p.Height)
		fmt.Printf("Weight: %d\n", p.Weight)
		fmt.Printf("Stats:\n")
		fmt.Printf("  -hp: %d\n", p.Stats[0].BaseStat)
		fmt.Printf("  -attack: %d\n", p.Stats[1].BaseStat)
		fmt.Printf("  -defense: %d\n", p.Stats[2].BaseStat)
		fmt.Printf("  -special-attack: %d\n", p.Stats[3].BaseStat)
		fmt.Printf("  -special-defense: %d\n", p.Stats[4].BaseStat)
		fmt.Printf("  -special-defense: %d\n", p.Stats[5].BaseStat)
		fmt.Printf("Types:\n")
		for _, t := range p.Types {
			fmt.Printf("  - %s\n", t.Type.Name)
		}

	default:
		fmt.Println("invalid command type")

	}

}
