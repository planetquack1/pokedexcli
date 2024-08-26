package main

import "fmt"

func createURL(cfg *config) string {

	return fmt.Sprintf("https://pokeapi.co/api/v2/%s/?limit=%d&offset=%d", cfg.endpoint, cfg.limit, cfg.limit*cfg.page)

}
