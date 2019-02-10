package data_test

import (
	"testing"

	"github.com/fairoldi/recipesapi/data"
)

var (
	testRecipe = data.Recipe{
		Title: "test recipe",
		Ingredients: []data.Ingredient{
			data.Ingredient{IngrName: "test ingredient", Qty: "test qty"},
		},
		Tools: []data.Tool{
			data.Tool{ToolName: "test tool", Notes: "test note"},
		},
		Steps: []data.Step{
			"test step 0",
			"test step 1",
		},
	}
)

const (
	correctUrl = "localhost:32769"
	wrongUrl   = "wrongurl:123"
)

func TestDBOk(t *testing.T) {

	data.Init(correctUrl)
	if err := data.DeleteAllRecipes(); err != nil {
		t.Error(err)
	}

	recs, err := data.Recipes()
	if err != nil {
		t.Error(err)
	}

	if len(recs) != 0 {
		t.Error("Database not empty")
	}

	id, err := data.NewRecipe(&testRecipe)
	if err != nil {
		t.Error(err)
	}

	if len(id) <= 0 {
		t.Error("Invalid generated id")
	}

	recs, err = data.Recipes()
	if err != nil {
		t.Error(err)
	}

	if len(recs) <= 0 {
		t.Error("Database empty")
	}

	rec, err := data.FindRecipeByID(id)
	if err != nil {
		t.Error(err)
	}

	if rec.ID != id {
		t.Error("ID not consistent")
	}

	if err := data.DeleteRecipe(id); err != nil {
		t.Error(err)
	}

	rec, err = data.FindRecipeByID(id)
	if err == nil {
		t.Error("Error expected")
	}
}

func TestDBNotOk(t *testing.T) {

	data.Init(wrongUrl)
	if err := data.DeleteAllRecipes(); err == nil {
		t.Error("Expected db error")
	}

	_, err := data.Recipes()
	if err == nil {
		t.Error("Expected db error")
	}

	id, err := data.NewRecipe(&testRecipe)
	if err == nil {
		t.Error("Expected db error")
	}

	_, err = data.FindRecipeByID(id)
	if err == nil {
		t.Error("Expected db error")
	}

	if err := data.DeleteRecipe(id); err == nil {
		t.Error("Expected db error")
	}

}
