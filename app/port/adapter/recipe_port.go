package adapter

import "backend-test/models"

type RecipeInputAdapter interface {
	InsertRecipe(input *models.Recipes) (output *models.Recipes, err error)
}

type RecipeOutputAdapter interface {
	InsertRecipe(input *models.Recipes) (output *models.Recipes, err error)
}
