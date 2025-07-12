package main

type cliCommand struct {
	name string
	description string
	callback func(location *config) error
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
	}
}

type config struct {
	Next string
	Previous string
}

var location = config{
	Next: "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
	Previous: "",
}
