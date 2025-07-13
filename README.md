# Pokedex CLI

Pokedex CLI is a command-line application written in Go that allows users to interactively explore Pokémon locations, catch Pokémon, and manage a personal Pokédex using data from the public [PokeAPI](https://pokeapi.co/). The app features a REPL (Read-Eval-Print Loop) interface, caching for API responses, and several commands for exploring and managing Pokémon data.

---

## Features

- **Explore Pokémon locations:** Browse paginated lists of location areas from the Pokémon world.
- **List Pokémon in an area:** See which Pokémon can be found in a specific location.
- **Catch Pokémon:** Attempt to catch Pokémon by name and add them to your personal Pokédex.
- **Inspect Pokémon:** View detailed stats and types for Pokémon you have caught.
- **View your Pokédex:** List all Pokémon you have caught so far.
- **Help and navigation:** Built-in help and navigation commands for ease of use.
- **Caching:** API responses are cached for 5 minutes to reduce redundant network requests.

---

## Installation

### Prerequisites

- Go 1.24.4 or later (see `go.mod`)

### Build

Clone the repository and build the CLI:

```sh
git clone https://github.com/frozendolphin/pokedexcli.git
cd pokedexcli
go build -o pokedexcli
```

### Run

Start the CLI:

```sh
./pokedexcli
```

You will see a prompt: `Pokedex >` where you can enter commands.

---

## Usage

Type `help` at the prompt to see all available commands:

- `help` — Displays a help message with all commands.
- `map` — Displays the next page of location areas.
- `mapb` — Displays the previous page of location areas.
- `explore <location>` — Lists all Pokémon in the specified location area.
- `catch <pokemon>` — Attempts to catch a Pokémon by name.
- `inspect <pokemon>` — Shows stats for a caught Pokémon.
- `pokedex` — Lists all Pokémon you have caught.
- `exit` — Exits the application.

Example session:

```
Pokedex > map
Pokedex > explore pallet-town
Pokedex > catch pikachu
Pokedex > pokedex
Pokedex > inspect pikachu
Pokedex > exit
```

---

## Configuration

- **No environment variables or configuration files are required.**
- The application uses the public PokeAPI and caches responses in memory for 5 minutes.

---

## Testing

To run the tests:

```sh
go test ./...
```

---

## Project Structure

- `main.go` — Entry point, starts the REPL.
- `repl.go` — REPL logic.
- `command_*.go` — Individual command implementations.
- `registry.go` — Command registration and configuration.
- `internal/pokecache/` — Simple in-memory cache for API responses.

---

## License

This project is open source and available under the MIT License.