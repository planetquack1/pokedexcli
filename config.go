package main

import "github.com/planetquack1/pokedexcli/internal/pokecache"

type Config struct {
	cache       *pokecache.Cache
	endpoint    string
	location    string
	commandType string
	limit       int
	page        int // starts at page -1 (no page)
}
