package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	result := strings.Fields(lower)
	return result
}

func startRepl() {
	prefix := "Pokedex > "
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("%s", prefix)
		text := scanner.Text()
		cleanText := cleanInput(text)
		if len(cleanText) == 0 {
			continue
		}

		command := cleanText[0]
		fmt.Printf("Your command was: %s\n", command)
	}
}
