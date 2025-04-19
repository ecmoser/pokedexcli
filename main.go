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
	callback    func(*config) error
}

type config struct {
	next     string
	previous string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 locations",
			callback:    CommandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations",
			callback:    CommandMapb,
		},
	}
}

func CleanInput(input string) []string {
	// Split input into words and convert to lowercase
	words := strings.Fields(strings.ToLower(input))
	return words
}

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	helpMsg := "Welcome to the Pokedex!\nUsage:\n\n"
	for _, command := range getCommands() {
		helpMsg += fmt.Sprintf("%s: %s\n", command.name, command.description)
	}
	fmt.Println(helpMsg)
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	initConfig := config{
		previous: "",
		next:     "https://pokeapi.co/api/v2/location-area?offset=0",
	}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleanText := CleanInput(text)
		handlerFound := false
		for _, command := range getCommands() {
			if len(cleanText) == 0 {
				continue
			}
			if command.name == cleanText[0] {
				if err := command.callback(&initConfig); err != nil {
					fmt.Println(err)
				}
				handlerFound = true
				break
			}
		}
		if !handlerFound {
			fmt.Println("Unknown command")
		}
	}
}
