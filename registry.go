package main

import (
	"time"

	"github.com/frozendolphin/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name string
	description string
	callback func(location *config, args []string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand {
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Displays a page of location areas",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Displays previous page of location areas",
			callback: commandMapb,
		},
		"explore": {
			name: "explore",
			description: "Displays a list of all the pokemon in the area",
			callback: commandExplore,
		},
		"catch": {
			name: "catch",
			description: "Attempt to catch a pokemon that is not in inventory",
			callback: commandCatch,
		},
		"inspect": {
			name: "inspect",
			description: "Gives pokemon's stats if in inventory",
			callback: commandInspect,
		},
	}
}

type config struct {
	Next string
	Previous string
	TheCache *pokecache.Cache
}

var location = config{
	Next: "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
	Previous: "",
	TheCache: pokecache.NewCache(5 * time.Minute),
}
