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
	callback		func() error
}

var commands map[string]cliCommand

func startRepl() {
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
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		scanedString := scanner.Text()
		wordSlices := cleanInput(scanedString)
		found := false
		for n,m := range commands{

			if wordSlices[0] == n{
				m.callback()
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