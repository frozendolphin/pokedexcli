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

func simpleREPL() error {
	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userinput := scanner.Text()
		cleaned_slice := cleanInput(userinput)
		if len(cleaned_slice) == 0 {
			continue
		}
		first_word := cleaned_slice[0]
		val, ok := getCommands()[first_word]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		var args []string
		if len(cleaned_slice) > 1 {
			args = cleaned_slice[1:]
		} else {
			args = nil
		}
		err := val.callback(&location, args)
		if err != nil {
			fmt.Println("Something went wrong...", err)
		}
	}
}



