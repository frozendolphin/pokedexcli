package main

import (
	"math/rand"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PokemonDetails struct {
	BaseExp int `json:"base_experience"`
	Name string `json:"name"`
	Height int `json:"height"`
	Weight int `json:"weight"`
	Stats []struct {
		Base_stat int `json:"base_stat"`
		Stat_name struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

var pokemon_caught = make(map[string]PokemonDetails)

func commandCatch(location *config, args []string) error {
	if args == nil {
		fmt.Println("Need to give pokemon name with catch command")
		return nil
	}
	pokemon_name := args[0] 

	url := "https://pokeapi.co/api/v2/pokemon/" + pokemon_name

	if _, ok := pokemon_caught[pokemon_name]; ok {
		fmt.Println("Pokemon already in inventory!")
		return nil
	}
	
	var data []byte
	val, ok := location.TheCache.Get(url)
	if ok {
		data = val
	} else {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}
	}
	
	var poke_details PokemonDetails

	if err := json.Unmarshal(data, &poke_details); err != nil {
		fmt.Println(pokemon_name, "doesn't exist!")
		return nil
	}

	caught := false
	if rand.Intn(100) < int(85.0 - (float64(poke_details.BaseExp)/15.0)) {
		caught = true
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", poke_details.Name,)
	if !caught {
		fmt.Println(poke_details.Name, "escaped!")
		return nil
	}
	fmt.Println(poke_details.Name, "was caught!")

	pokemon_caught[pokemon_name] = poke_details
	if !ok {
		location.TheCache.Add(pokemon_name, data)
	}

	return nil
}