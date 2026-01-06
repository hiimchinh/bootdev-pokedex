package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Showing help",
			callback:    commandHelp,
		},
	}
	for {
		fmt.Print("Pokedex > ")
		for scanner.Scan() {
			line := scanner.Text()
			arr := strings.Fields(line)
			firstWord := strings.ToLower(arr[0])
			command, ok := commands[firstWord]
			if !ok {
				fmt.Println("Unknown command")
				break
			}
			command.callback()
			break
		}

	}
}
