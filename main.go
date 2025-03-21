package main

import (
	"time"

	"github.com/BambiCPT/pokedexcli/commands"
	"github.com/BambiCPT/pokedexcli/internal/pokeapi"
	"github.com/BambiCPT/pokedexcli/repl"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &commands.Config{
		Client: pokeClient,
	}

	repl.StartRepl(cfg)
}
