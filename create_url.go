package main

import (
	"fmt"
)

func createURL(cfg *Config) string {

	switch cfg.commandType {
	case "map":
		return fmt.Sprintf("https://pokeapi.co/api/v2/%s/?limit=%d&offset=%d", cfg.endpoint, cfg.limit, cfg.limit*cfg.page)
	case "explore":
		return fmt.Sprintf("https://pokeapi.co/api/v2/%s/%s/", cfg.endpoint, cfg.location)
	default:
		return "unknown command type"
	}

}
