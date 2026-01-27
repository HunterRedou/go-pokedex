package main

import(
	"net/http"
	"github.com/HunterRedou/pokedex/internal/pokeapi"
	"time"
)



func main(){
cfg := &config{
    pokeapi: pokeapi.NewClient(&http.Client{}, 5*time.Minute),
		caughtPokemon: make(map[string]pokeapi.Catch),
	}
	startRepl(cfg)
}
