package main

import "github.com/planetquack1/pokedexcli/internal/pokecache"

type config struct {
	cache    *pokecache.Cache
	endpoint string
	limit    int
	page     int // starts at page -1 (no page)
}
