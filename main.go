package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/hiimchinh/bootdev-pokedex/cmd"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *cmd.Config) error
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    cmd.CommandExit,
		},
		"help": {
			name:        "help",
			description: "Showing help",
			callback:    cmd.CommandHelp,
		},
		"map": {
			name:        "map",
			description: "list location areas in pokemon world",
			callback:    cmd.CommandMap,
		},
		"mapb": {
			name:        "map back",
			description: "list location areas in previous page",
			callback:    cmd.CommandMapb,
		},
	}
	nextUrl := "https://pokeapi.co/api/v2/location-area"
	cfg := &cmd.Config{
		Next:     &nextUrl,
		Previous: nil,
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
			command.callback(cfg)
			break
		}

	}
}
