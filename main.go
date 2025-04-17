package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CleanInput(input string) []string {
	// Split input into words and convert to lowercase
	words := strings.Fields(strings.ToLower(input))
	return words
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleanText := CleanInput(text)
		fmt.Printf("Your command was: %v\n", cleanText[0])
	}
}
