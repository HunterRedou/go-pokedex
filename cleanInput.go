package main

import(
	"strings"
	"bufio"
	"os"
	"fmt"
)

type cliCommand struct{
	name			string
	description 	string
	callback		func(*config, ...string) error
}

var commands map[string]cliCommand

func startRepl(cfg *config) {
	commands = map[string]cliCommand{
		"exit": {
			name: 			"exit",
			description: 	"Exit the Pokedex",
			callback: 		commandExit,
		},
		"help": {
			name:			"help",
			description:	"Displays a help message",
			callback: 		commandHelp,
		},
		"map": {
			name:			"map",
			description:	"Shows Area Locations",
			callback:		commandMap,
		},
		"mapb": {
			name:			"mapb",
			description:	"Goes back to the Previous Locations",
			callback:		commandMapb,
		},
		"explore": {
			name:			"explore",
			description:	"Explores the Area Location",
			callback:		commandExplore,
		},
		"catch": {
			name:			"catch",
			description:	"Catch a Pokemon",
			callback:		commandCatch,
		},
		"inspect": {
			name:			"inspect",
			description:	"Inspect the Info over Catched Pokemon",
			callback:		commandInspect,
		},
		"pokedex": {
			name:			"pokedex",
			description:	"Shows Pokedex with your caught Pokemon",
			callback:		commandPokedex,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		scanedString := scanner.Text()
		wordSlices := cleanInput(scanedString)
		found := false
		args := []string{}
		if len(wordSlices) > 1{
			args = wordSlices[1:]
		}
		for n,m := range commands{
			
			if wordSlices[0] == n{
				m.callback(cfg, args...)
				found =true
			}
		}
		if !found{
			fmt.Print("Unkown command")
		}
		
	}
}


func cleanInput(text string) []string{
	low := strings.ToLower(text)
	removedWs:= strings.Fields(low)
	return removedWs

}
