package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mtslzr/pokeapi-go"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func matchCommand(input string, cfg *Config) cliCommand {

	// Check if the input starts with "explore" or "catch"
	if strings.HasPrefix(input, "explore ") {
		location := strings.TrimPrefix(input, "explore ")
		return cliCommand{
			name:        "explore",
			description: "Explore the specified location",
			callback: func() error {
				return cfg.commandExplore(location)
			},
		}
	}
	if strings.HasPrefix(input, "catch ") {
		pokemonToCatch := strings.TrimPrefix(input, "catch ")
		return cliCommand{
			name:        "catch",
			description: "Catch the specified Pokemon",
			callback: func() error {
				return cfg.commandCatch(pokemonToCatch)
			},
		}
	}

	// Handle other commands
	switch input {
	case "help":
		return cliCommand{
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		}
	case "exit":
		return cliCommand{
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		}
	case "map":
		return cliCommand{
			name:        "map",
			description: "Get the next 20 locations",
			callback:    cfg.commandMap,
		}
	case "mapb":
		return cliCommand{
			name:        "mapb",
			description: "Get the previous 20 locations",
			callback:    cfg.commandMapb,
		}
	case "pokedex":
		return cliCommand{
			name:        "pokedex",
			description: "List the pokedex",
			callback:    cfg.commandPokedex,
		}
	default:
		return cliCommand{
			name:        "unknown",
			description: "Try another command",
			callback:    commandUnknown,
		}
	}
}

func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\n")
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}

func (cfg *Config) commandExplore(location string) error {
	cfg.commandType = "explore"
	cfg.endpoint = "location-area"

	cfg.location = location

	fmt.Printf("Exploring %s area...\n", location)

	getAndPrintResponse(cfg)

	return nil
}

func (cfg *Config) commandCatch(pokemonToCatch string) error {
	cfg.commandType = "catch"
	cfg.endpoint = "pokemon"

	p, err := pokeapi.Pokemon(pokemonToCatch)
	if err != nil {
		fmt.Printf("Error retrieving %s\n", pokemonToCatch)
		return nil
	}

	cfg.pokemon = p
	cfg.baseExperience = p.BaseExperience

	fmt.Printf("Pokemon has base experience %d\n", p.BaseExperience)

	getAndPrintResponse(cfg)

	return nil
}

func (cfg *Config) commandMap() error {
	cfg.commandType = "map"
	cfg.endpoint = "location"
	cfg.page++

	getAndPrintResponse(cfg)

	return nil
}

func (cfg *Config) commandMapb() error {
	cfg.commandType = "map"
	cfg.endpoint = "location"
	if cfg.page <= 0 {
		fmt.Printf("Warning: you are on the first page.\n")
		return nil
	}
	cfg.page--

	getAndPrintResponse(cfg)

	return nil
}

func (cfg *Config) commandPokedex() error {

	fmt.Println("Your Pokedex:")

	for _, pokemon := range cfg.pokedex {
		fmt.Printf("- %s\n", pokemon.Name)
	}

	return nil
}

func commandUnknown() error {
	fmt.Printf("Unkown command. Try again!\n")
	return nil
}
