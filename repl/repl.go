package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/BambiCPT/pokedexcli/commands"
)

func StartRepl(cfg *commands.Config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}

		commandName := input[0]
		command, exists := commands.GetCommands()[commandName]
		if exists {
			err := command.Callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	loweredText := strings.ToLower(text)
	cleanOutput := strings.Fields(loweredText)
	return cleanOutput
}
