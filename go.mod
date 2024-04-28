module github.com/crayboi420/pokedexcli 

go 1.22.2

replace github.com/crayboi420/pokedexcli/internal/pokecache => ./internal/pokecache

require (
	github.com/crayboi420/pokedexcli/internal/pokecache v0.0.0
)