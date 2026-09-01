package main

import "github.com/silentfin/pokedex/internal/pokeapi"

func main() {
	configs := config{commands: getCommands(), inventory: map[string]pokeapi.Pokemon{}}
	repl(&configs)
}
