package repository

import (
	"pokemon/api/pokemon/domain/valueobject"
)

type PokemonRepository interface {
	GetPokemons(limit string, offset string) (pokemons valueobject.PokemonListResponse, err error)
	GetPokemonByID(id string) (pokemon valueobject.PokemonResponse, err error)
	GetPokemonByNombre(nombre string) (pokemon valueobject.PokemonResponse, err error)
}
