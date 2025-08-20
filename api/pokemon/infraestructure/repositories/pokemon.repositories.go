package repositories

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"pokemon/api/pokemon/domain/repository"
	"pokemon/api/pokemon/domain/valueobject"
	"pokemon/utils"
)

type PokemonRepository struct {
}

func (p PokemonRepository) GetPokemons(limit string, offset string) (pokemons valueobject.PokemonListResponse, err error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon?limit=%s&offset=%s", limit, offset)
	body, err, _ := utils.MakeRequest(url)
	if err != nil {
		return valueobject.PokemonListResponse{}, fmt.Errorf("error de peticion http(Pokemon): %w", err)
	}

	// Parseaamos el JSON a la estructura
	err = json.Unmarshal(body, &pokemons)
	if err != nil {
		return valueobject.PokemonListResponse{}, fmt.Errorf("error parsing JSON: %w", err)
	}

	return pokemons, nil
}

func (p PokemonRepository) GetPokemonByID(id string) (pokemon valueobject.PokemonResponse, err error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", id)
	body, err, statusCode := utils.MakeRequest(url)
	if err != nil {
		return valueobject.PokemonResponse{}, fmt.Errorf("error de peticion http(Pokemon): %w", err)
	}
	if statusCode == 404 {
		return valueobject.PokemonResponse{}, fmt.Errorf("Pokemon no encontrado")
	}

	// Parseaamos el JSON a la estructura
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return valueobject.PokemonResponse{}, fmt.Errorf("error parsing JSON: %w", err)
	}

	return pokemon, nil
}

func (p PokemonRepository) GetPokemonByNombre(nombre string) (pokemon valueobject.PokemonResponse, err error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", nombre)
	body, err, statusCode := utils.MakeRequest(url)
	if err != nil {
		return valueobject.PokemonResponse{}, fmt.Errorf("error de peticion http(Pokemon): %w", err)
	}
	if statusCode == 404 {
		return valueobject.PokemonResponse{}, fmt.Errorf("Pokemon no encontrado")
	}
	// Parseaamos el JSON a la estructura
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		log.Error("Error en ", err)
		return valueobject.PokemonResponse{}, fmt.Errorf("error parsing JSON: %w", err)
	}

	return pokemon, nil
}
func NewPokemonRepository() repository.PokemonRepository {
	return PokemonRepository{}
}
