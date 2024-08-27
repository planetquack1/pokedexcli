package main

import (
	"github.com/mtslzr/pokeapi-go/structs"
	"github.com/planetquack1/pokedexcli/internal/pokecache"
)

type Config struct {
	cache          *pokecache.Cache
	pokedex        map[string]structs.Pokemon
	endpoint       string
	location       string
	pokemon        structs.Pokemon
	commandType    string
	baseExperience int
	limit          int
	page           int // starts at page -1 (no page)
}
