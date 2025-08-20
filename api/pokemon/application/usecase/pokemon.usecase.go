package usecase

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"pokemon/api/pokemon/domain/repository"
	"pokemon/api/pokemon/domain/usecase"
	"pokemon/api/pokemon/domain/valueobject"
	"time"
)

type PokemonUseCase struct {
	repo repository.PokemonRepository
}

func (p PokemonUseCase) GetPokemonsFoto(limit string, offset string) (pokemonsFoto []valueobject.PokemonItemFotoResponse, error error) {
	var result []valueobject.PokemonItemFotoResponse

	pokemons, error := p.repo.GetPokemons(limit, offset)
	if error != nil {
		return pokemonsFoto, nil
	}

	for _, pokemon := range pokemons.Results {
		pok, error := p.repo.GetPokemonByNombre(pokemon.Name)
		if error != nil {
			fmt.Printf("Error obteniendo %s: %v\n", pok.Name, error)
			continue
		}

		result = append(result, valueobject.PokemonItemFotoResponse{
			Id:              pok.ID,
			Name:            pokemon.Name,
			FotoURL:         pok.Sprites.FrontDefault,
			Height:          pok.Height,
			Weight:          pok.Weight,
			BaseExperiencia: pok.BaseExperience,
			Types:           pok.Types,
			Habilidades:     pok.Abilities,
		})

		// Pequeña pausa para no saturar la API
		time.Sleep(100 * time.Millisecond)

	}

	return result, error
}

func (p PokemonUseCase) GetPokemons(limit string, offset string) (pokemons valueobject.PokemonListResponse, err error) {
	pokemons, err = p.repo.GetPokemons(limit, offset)

	return
}

func (p PokemonUseCase) GetPokemonByID(id string) (pokemon valueobject.PokemonResponse, error error) {

	pokemon, error = p.repo.GetPokemonByID(id)

	return
}

func (p PokemonUseCase) GetPokemonByName(nombrePokemon string) (pokemon valueobject.PokemonItemFotoResponse, error error) {
	dataPokemon, error := p.repo.GetPokemonByNombre(nombrePokemon)
	if error != nil {
		log.Error("Error al obtener pokemon", error)
		return
	}
	//Formateamos los datos
	pokemon = valueobject.PokemonItemFotoResponse{
		Id:              dataPokemon.ID,
		Name:            dataPokemon.Name,
		FotoURL:         dataPokemon.Sprites.FrontDefault,
		Height:          dataPokemon.Height,
		Weight:          dataPokemon.Weight,
		BaseExperiencia: dataPokemon.BaseExperience,
		Types:           dataPokemon.Types,
		Habilidades:     dataPokemon.Abilities,
	}

	return
}

func NewPokemonUseaCase(rep repository.PokemonRepository) usecase.PokemonUseCase {
	return &PokemonUseCase{
		repo: rep,
	}
}
