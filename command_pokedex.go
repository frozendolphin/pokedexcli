package main

import "fmt"

func commandPokedex(location *config, args []string) error {
	if len(pokemon_caught) == 0 {
		fmt.Println("You have no pokemon in your pokedex :(")
		return nil
	}
	fmt.Println("Your Pokedex")
	for key, _ := range pokemon_caught {
		fmt.Printf("  - %v\n", key)
	}
	return nil
}