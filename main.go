package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/planetquack1/pokedexcli/internal/pokecache"
)

func main() {

	c := pokecache.NewCache(5 * time.Second)

	cfg := Config{
		cache:       &c,
		endpoint:    "location",
		location:    "",
		commandType: "map",
		limit:       20,
		page:        -1,
	}

	fmt.Printf("Pokedex > ")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		cmd := matchCommand(scanner.Text(), &cfg)
		cmd.callback()

		fmt.Printf("Pokedex > ")
	}
}
