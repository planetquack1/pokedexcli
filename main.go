package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	cfg := config{
		resource: "locations",
		limit:    20,
		page:     0,
	}

	fmt.Printf("Pokedex > ")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		cmd := matchCommand(scanner.Text(), &cfg)

		cmd.callback()

		fmt.Printf("Pokedex > ")
	}
}
