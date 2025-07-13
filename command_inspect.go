package main

import "fmt"

func commandInspect(location *config, args []string) error {
	if args == nil {
		fmt.Println("Need to give pokemon name with inspect command")
		return nil
	}

	val, ok := pokemon_caught[args[0]]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %v\n", val.Name)
	fmt.Printf("Height: %v\n", val.Height)
	fmt.Printf("Weight: %v\n", val.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range val.Stats {
		fmt.Printf("  -%v: %v\n", stat.Stat_name.Name, stat.Base_stat)
	}
	fmt.Printf("Types:\n")
	for _, stat := range val.Types {
		fmt.Printf("  - %v\n", stat.Type.Name)
	}

	return nil
}