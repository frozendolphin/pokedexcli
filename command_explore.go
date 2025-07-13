package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PokemonsInArea struct{
	Pokemon_encounters []PokemonMetaInfo `json:"pokemon_encounters"`
}

type PokemonMetaInfo struct {
	Pokemon struct {
		Name string `json:"name"`
		Url string `json:"url"`
	} `json:"pokemon"`
}

func commandExplore(location *config, args []string) error {
	if args == nil {
		fmt.Println("Need to enter an argument with explore")
		return nil
	}
	city := args[0]
	fmt.Printf("Exploring %v...", city)

	url := "https://pokeapi.co/api/v2/location-area/" + city

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

	var pokelist PokemonsInArea

	if err := json.Unmarshal(data, &pokelist); err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, m_info := range pokelist.Pokemon_encounters {
		fmt.Println("-", m_info.Pokemon.Name)
	}

	if !ok {
		location.TheCache.Add(url, data)
	}

	return nil
}