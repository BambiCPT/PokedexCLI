package commands

import "github.com/BambiCPT/pokedexcli/internal/pokeapi"

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*Config) error
}

type Config struct {
	NextURL *string
	PrevURL *string
	Client  *pokeapi.Client
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    CommandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Explore the next 20 locations",
			Callback:    MapCommand,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Go back to the previous 20 locations",
			Callback:    MapBackCommand,
		},
	}
}
