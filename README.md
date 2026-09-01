# pokedex
A command-line Pokedex built in Go, using the [PokeAPI](https://pokeapi.co/) as the source of Pokémon data.

## Features
- Interactive REPL (Read-Eval-Print Loop) interface
- Paginated browsing of map/location areas
- Explore locations to see which Pokémon live there
- Catch Pokémon with a chance-based mechanic
- Inspect stats and details of Pokémon you've already caught
- Local caching of API responses for faster repeat lookups

## Getting Started
Clone the repository and build the binary:
 
```bash
git clone https://github.com/silentfin/pokedex.git
cd pokedex
go build -o pokedex
```
 
Then run it:
 
```bash
./pokedex
```
 
Alternatively, run it directly without building:
 
```bash
go run .
```


## Usage 
Once running, you'll land in the Pokedex REPL:
 
```
Pokedex > help
```
 
## Commands
 
| Command | Description |
|---|---|
| `help` | Displays a help message |
| `map` | Shows the next 20 map areas |
| `mapb` | Shows the previous 20 map areas |
| `explore <area-name>` | Lists all Pokémon found in the given area |
| `catch <pokemon-name>` | Attempts to catch the given Pokémon |
| `inspect <pokemon-name>` | Shows details for a Pokémon you've already caught |
| `pokedex` | Prints the list of all Pokémon you've caught |
| `exit` | Exits the Pokedex |

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
