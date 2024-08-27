package main

func addPokemon(cfg *Config) {
	cfg.pokedex[cfg.pokemon.Name] = cfg.pokemon
}
