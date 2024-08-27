package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
	"github.com/planetquack1/pokedexcli/internal/pokecache"
)

func main() {

	c := pokecache.NewCache(5 * time.Second)

	p, err := pokeapi.Pokemon("1")
	if err != nil {
		fmt.Println("Error connecting to PokeAPI server")
	}

	cfg := Config{
		cache:          &c,
		pokedex:        make(map[string]structs.Pokemon),
		endpoint:       "location",
		location:       "",
		pokemon:        p,
		commandType:    "map",
		baseExperience: 200, // default
		limit:          20,
		page:           -1,
	}

	fmt.Printf("Pokedex > ")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		cmd := matchCommand(scanner.Text(), &cfg)
		cmd.callback()

		fmt.Printf("Pokedex > ")
	}
}
