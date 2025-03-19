package main

import (
	"time"

	"github.com/BambiCPT/pokedexcli/commands"
	"github.com/BambiCPT/pokedexcli/internal/pokecache"
	"github.com/BambiCPT/pokedexcli/repl"
)

func main() {
	cache := pokecache.NewCache(5 * time.Second)

	cfg := commands.Config{
		NextURL: "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
		PrevURL: "",
		Cache:   cache,
	}

	repl.StartRepl(&cfg)
}
