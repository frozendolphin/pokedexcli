package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Location_area struct{
	Count int `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []Location_area_name `json:"results"`
}

type Location_area_name struct{
	Name string `json:"name"`
	Url string `json:"url"`
}

func location_list(url string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var areas Location_area

	if err:= json.Unmarshal(data, &areas); err != nil {
		return err
	}

	location.Next = areas.Next
	location.Previous = areas.Previous
	for _, item := range(areas.Results){
		fmt.Println(item.Name)
	}
	return  nil
}

func commandMap(location *config) error {
	url := location.Next
	
	err := location_list(url)
	if err != nil {
		return err
	}

	return nil
}

func commandMapb(location *config) error {
	url := location.Previous

	if url == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	
	err := location_list(url)
	if err != nil {
		return err
	}

	return nil
}