package main

import (
	"fmt"
)

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage: \n\n")
	for key, val := range getCommands() {
		fmt.Printf("%v: %v\n", key, val.description)
	}
	return nil
}