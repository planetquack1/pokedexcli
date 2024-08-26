package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func matchCommand(input string, cfg *config) cliCommand {

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

func (cfg *config) commandMap() error {
	cfg.page++

	getAndPrintPage(cfg)

	return nil
}

func (cfg *config) commandMapb() error {
	if cfg.page <= 0 {
		fmt.Printf("Error: you are on the first page.\n")
		return nil
	}
	cfg.page--

	getAndPrintPage(cfg)

	return nil
}

func commandUnknown() error {
	fmt.Printf("Unkown command. Try again!\n")
	return nil
}
