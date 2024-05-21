module github.com/kanavj/pokedexcli 

go 1.22.2

replace github.com/kanavj/pokedexcli/internal/pokecache => ./internal/pokecache

require (
	github.com/kanavj/pokedexcli/internal/pokecache v0.0.0
)