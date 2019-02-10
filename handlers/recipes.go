package handlers

import (
	"net/http"

	"github.com/fairoldi/recipesapi/data"

	"github.com/labstack/echo"
)


func Init(url string) {
	data.Init(url)
}

// Recipes is the handler to retrieve all recipes
func Recipes(c echo.Context) error {

	recipes, err := data.Recipes()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if len(recipes) == 0 {
		recipes = []data.Recipe{}
	}
	return c.JSON(http.StatusOK, recipes)
}

// NewRecipe creates a new recipe in the database
func NewRecipe(c echo.Context) error {

	recp := &data.Recipe{}
	if err := c.Bind(recp); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if _, err := data.NewRecipe(recp); err != nil {
		return echo.NewHTTPError(http.StatusConflict, err.Error())
	}

	return c.JSON(http.StatusCreated, *recp)
}

// FindRecipeByID retrieves a recipe by ID
func FindRecipeByID(c echo.Context) error {

	recp, err := data.FindRecipeByID(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, *recp)
}

// DeleteRecipe deletes a recipe 
func DeleteRecipe(c echo.Context) error {

	err := data.DeleteRecipe(c.Param("id"))
	if (err != nil) {
		return echo.NewHTTPError(http.StatusNotFound, err.Error)
	}

	return c.NoContent(http.StatusOK)

}
