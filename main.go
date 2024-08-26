package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	cfg := config{
		endpoint: "locations",
		limit:    20,
		page:     0,
	}

	fmt.Printf("Pokedex > ")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		cmd := matchCommand(scanner.Text(), &cfg)
		cmd.callback()

		url := createURL(&cfg)

		fmt.Printf("URL: %s\n", url)
		fmt.Printf("Pokedex > ")
	}
}
