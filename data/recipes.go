package data

import (
	"fmt"

	r "gopkg.in/rethinkdb/rethinkdb-go.v5"
)

var (
	session *r.Session
)

// Ingredient data type
type Ingredient struct {
	IngrName string `rethinkdb:"ingredient"`
	Qty      string `rethinkdb:"qty"`
}

// Tool data type
type Tool struct {
	ToolName string `rethinkdb:"tool"`
	Notes    string `rethinkdb:"notes"`
}

// Step data type
type Step string

// Recipe data type
type Recipe struct {
	ID          string       `rethinkdb:"id,omitempty"`
	Title       string       `rethinkdb:"title"`
	Ingredients []Ingredient `rethinkdb:"ingredients"`
	Tools       []Tool       `rethinkdb:"tools"`
	Steps       []Step       `rethinkdb:"steps"`
}

// Init initializes a connection to the database
func Init(url string) {
	var err error
	session, err = r.Connect(r.ConnectOpts{
		Address:  url,
		Database: "recipes",
	})

	if err != nil {
		fmt.Println(err)
	}
}

// Recipes retrieves all recipes from RethinkDB datastore.
func Recipes() ([]Recipe, error) {

	res, err := r.Table("recipes").Run(session)
	defer res.Close()
	if err != nil {
		return nil, err
	}

	var recipes []Recipe
	if err := res.All(&recipes); err != nil {
		return nil, err
	}

	return recipes, nil
}

// NewRecipe stores a new recipe in the database
func NewRecipe(recp *Recipe) (string, error) {

	result, err := r.Table("recipes").Insert(*recp).RunWrite(session)
	if err != nil {
		return "", err
	}

	recp.ID = result.GeneratedKeys[0]

	return result.GeneratedKeys[0], nil
}

// FindRecipeByID retrieves a recipe by ID
func FindRecipeByID(id string) (*Recipe, error) {

	res, err := r.Table("recipes").Get(id).Run(session)
	defer res.Close()
	if err != nil {
		return nil, err
	}

	var recipe Recipe

	if err := res.One(&recipe); err != nil {
		return nil, err
	}

	return &recipe, nil
}

// DeleteRecipe deletes a recipe from the database
func DeleteRecipe(id string) error {

	_, err := r.Table("recipes").Get(id).Delete().RunWrite(session)
	if err != nil {
		return err
	}

	return nil
}

// DeleteAllRecipes deletes all recipes from the database
func DeleteAllRecipes() error {
	_, err := r.Table("recipes").Delete().RunWrite(session)
	if err != nil {
		return err
	}

	return nil
}

/*
{

    "id": "224a8edf-7bd8-4ec1-bb11-06145f3f64ba" ,
    "ingredients": [
        {
            "ingrName": "bacon" ,
            "qty": "4 slices"
        } ,
        {
            "ingrName": "eggs" ,
            "qty": "2"
        } ,
        {
            "ingrName": "butter" ,
            "qty": "1 teaspoon"
        }
    ] ,
    "steps": [
        "Put bacon strips in the cold skillet" ,
        "Place the skillet on medium heat" ,
        "After about 3 minutes, flip the bacon" ,
        "after another 3 minutes, take the bacon strips off the skillet" ,
        "with some paper towels, dry up some of the fat" ,
        "break the eggs" ,
        "put the eggs in the skillet" ,
        "cover the skillet with a lid" ,
        "turn off the fire" ,
        "wait 3 minutes"
    ] ,
    "title": "Bacon and eggs" ,
    "tools": [
        {
            "notes": "12"" ,
            "toolName": "skillet"
        } ,
        {
            "notes": "" ,
            "toolName": "spatula"
        }
    ]

}
*/
