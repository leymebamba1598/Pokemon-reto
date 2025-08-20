package api

import (
	"github.com/gin-gonic/gin"
	"pokemon/api/pokemon/infraestructure/http/routes"
)

func RegisterRoutes(e *gin.Engine) {
	routes.PokemonRoutes(e)
}
