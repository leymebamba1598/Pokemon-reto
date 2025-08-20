package usecase

import (
	"pokemon/api/pokemon/domain/valueobject"
)

type PokemonUseCase interface {
	GetPokemons(limit string, offset string) (pokemons valueobject.PokemonListResponse, err error)
	GetPokemonsFoto(limit string, offset string) (pokemons []valueobject.PokemonItemFotoResponse, err error)
	GetPokemonByID(id string) (pokemon valueobject.PokemonResponse, error error)
	GetPokemonByName(nombrePokemon string) (pokemon valueobject.PokemonItemFotoResponse, error error)
}
