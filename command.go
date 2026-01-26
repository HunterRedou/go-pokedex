package main

import(
	"fmt"
	"os"
	"github.com/HunterRedou/pokedex/internal/pokeapi"
)

type config struct{
	pokeapi *pokeapi.Client
	next *string
	prev *string
}

func commandExit(cfg *config) error{
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error{
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range commands{
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(cfg *config) error{
	loc, err := cfg.pokeapi.GetBody(cfg.next)
	if err != nil{
		return err
	}

	cfg.next = loc.Next 
	cfg.prev = loc.Previous


	pokeapi.GetNames(loc)
	return nil
}

func commandMapb(cfg *config) error{

	if cfg.prev == nil{
		fmt.Println("Your on the first Page")
	}
	loc, err := cfg.pokeapi.GetBody(cfg.prev)
	if err != nil{
		return err
	}
	

	cfg.next = loc.Next
	cfg.prev = loc.Previous

	pokeapi.GetNames(loc)
	return nil
}
