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
	case "next":
		return cliCommand{
			name:        "next",
			description: "Get the next 20 locations",
			callback:    cfg.commandNext,
		}
	case "back":
		return cliCommand{
			name:        "next",
			description: "Get the previous 20 locations",
			callback:    cfg.commandBack,
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

func (cfg *config) commandNext() error {
	cfg.page++
	fmt.Printf("Page number: %d\n", cfg.page)
	return nil
}

func (cfg *config) commandBack() error {
	if cfg.page == 0 {
		fmt.Printf("Error: you are on the first page.\n")
		return nil
	}
	cfg.page--
	fmt.Printf("Page number: %d\n", cfg.page)
	return nil
}

func commandUnknown() error {
	fmt.Printf("Unkown command. Try again!\n")
	return nil
}
