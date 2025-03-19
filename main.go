package main

import (
	"github.com/BambiCPT/pokedexcli/commands"
	"github.com/BambiCPT/pokedexcli/repl"
)

func main() {
	cfg := commands.Config{
		NextURL: "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
		PrevURL: "",
	}

	repl.StartRepl(&cfg)
}
