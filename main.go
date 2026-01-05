package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		for scanner.Scan() {
			line := scanner.Text()
			arr := strings.Fields(line)
			firstWord := strings.ToLower(arr[0])
			fmt.Printf("Your command was: %v\n", firstWord)
			break

		}

	}
}
