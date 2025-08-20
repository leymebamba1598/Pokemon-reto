package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pokemon/api/pokemon/domain/usecase"
)

type PokemonController struct {
	useCase usecase.PokemonUseCase
}

func (p *PokemonController) GetPokemons(c *gin.Context) {
	limit := c.DefaultQuery("limit", "10")
	offset := c.DefaultQuery("offset", "0")
	pokemons, err := p.useCase.GetPokemons(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, pokemons)

}

func (p *PokemonController) GetPokemonsFoto(c *gin.Context) {
	limit := c.DefaultQuery("limit", "10")
	offset := c.DefaultQuery("offset", "0")
	pokemons, err := p.useCase.GetPokemonsFoto(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, pokemons)

}

func (p *PokemonController) GetPokemonByID(c *gin.Context) {
	id := c.Param("id")

	pokemon, err := p.useCase.GetPokemonByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, pokemon)

}

func (p *PokemonController) GetPokemonByNombre(c *gin.Context) {
	nombre := c.Param("name")

	pokemon, err := p.useCase.GetPokemonByName(nombre)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, pokemon)

}

func NewPokemonController(usecase usecase.PokemonUseCase) *PokemonController {
	return &PokemonController{
		useCase: usecase,
	}
}
