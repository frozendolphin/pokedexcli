package main

import (
	"strings"
	"os"
	"bufio"
	"fmt"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func simpleREPL() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			userinput := scanner.Text()
			cleaned_slice := cleanInput(userinput)
			if len(cleaned_slice) == 0 {
				continue
			}
			first_word := cleaned_slice[0]
			fmt.Println("Your command was:", first_word)
		}
	}
}