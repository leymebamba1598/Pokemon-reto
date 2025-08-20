package routes

import (
	"github.com/gin-gonic/gin"
	"pokemon/api/pokemon/application/usecase"
	"pokemon/api/pokemon/infraestructure/http/controller"
	"pokemon/api/pokemon/infraestructure/repositories"
)

func PokemonRoutes(e *gin.Engine) {
	repo := repositories.NewPokemonRepository()
	useCase := usecase.NewPokemonUseaCase(repo)
	controller := controller.NewPokemonController(useCase)

	e.GET("/api/v1/pokemon", controller.GetPokemons)
	e.GET("/api/v1/pokemon/foto", controller.GetPokemonsFoto)
	e.GET("/api/v1/pokemon/id/:id", controller.GetPokemonByID)
	e.GET("/api/v1/pokemon/name/:name", controller.GetPokemonByNombre)
}
