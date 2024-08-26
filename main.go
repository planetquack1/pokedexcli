package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	cfg := config{
		endpoint: "location",
		limit:    20,
		page:     -1,
	}

	fmt.Printf("Pokedex > ")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		cmd := matchCommand(scanner.Text(), &cfg)
		cmd.callback()

		fmt.Printf("Pokedex > ")
	}
}
