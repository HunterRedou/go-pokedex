package main

import(
	"net/http"
	"github.com/HunterRedou/pokedex/internal/pokeapi"
)



func main(){
cfg := &config{
    pokeapi: pokeapi.NewClient(&http.Client{}),
	}
	startRepl(cfg)
}
