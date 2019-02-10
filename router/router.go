package router

import (
	"github.com/fairoldi/recipesapi/handlers"
	"github.com/labstack/echo"
)

// InitRoutes initializes all routes for the recipes api
func InitRoutes(ep *echo.Echo, url string) {
	handlers.Init(url)
	ep.GET("/recipes", handlers.Recipes)
	ep.POST("/recipes", handlers.NewRecipe)
	ep.GET("/recipes/:id", handlers.FindRecipeByID)
	ep.DELETE("/recipes/:id", handlers.DeleteRecipe)
}
