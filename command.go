package main

import(
	"fmt"
	"os"
	"math/rand"
	"github.com/HunterRedou/pokedex/internal/pokeapi"
)

type config struct{
	pokeapi *pokeapi.Client
	next *string
	prev *string
	caughtPokemon map[string]pokeapi.Catch
}

func commandExit(cfg *config, args ...string) error{
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, args ...string) error{
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range commands{
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(cfg *config, args ...string) error{
	loc, err := cfg.pokeapi.GetBody(cfg.next)
	if err != nil{
		return err
	}

	cfg.next = loc.Next 
	cfg.prev = loc.Previous


	pokeapi.GetNames(loc)
	return nil
}

func commandMapb(cfg *config, args ...string) error{

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

func commandExplore(cfg *config, args ...string) error{
	if len(args) != 1{
		fmt.Println("Locations not in Command")
		return nil
	}
	names, err := cfg.pokeapi.GetPokemon(args[0])
	if err != nil{
		return err
	}

	fmt.Printf("Exploring %s...\n", args)
	fmt.Println("Found Pokemon:")

	for _, value := range names.Encounter{
		fmt.Println(" - " + value.Pokemon.Name)
	}

	return nil

}

func commandCatch(cfg *config, args ...string) error{
	if len(args) != 1{
		fmt.Print("The Pokemon is not in the Command")
		return nil
	}

	name := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	
	poke, err := cfg.pokeapi.CatchPokemon(name)
	if err != nil{
		return err
	}
	
	catched := rand.Intn(poke.BaseExp)
	if catched < 25{
		fmt.Printf("%s was caught!\n", name)
		cfg.caughtPokemon[poke.Name] = poke
		return nil
	}

	fmt.Printf("%s escaped!\n", name)
	return nil
	
}
