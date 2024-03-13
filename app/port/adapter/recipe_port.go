package adapter

import "backend-test/models"

type RecipeInputAdapter interface {
	GetListRecipe(request *models.RequestListRecipe) (output *models.ResponseListRecipe, err error)
	InsertRecipe(input *models.Recipes) (output *models.Recipes, err error)
	UpdateRecipe(id int, input *models.Recipes) (output *models.Recipes, err error)
	DeleteRecipe(id int) (err error)
}

type RecipeOutputAdapter interface {
	ListRecipe(req *models.RequestListRecipe) (output []*models.Recipes, count int64, err error)
	GetRecipeByID(id int) (output *models.Recipes, err error)
	InsertRecipe(input *models.Recipes) (output *models.Recipes, err error)
	UpdateRecipe(id int, input *models.Recipes) (output *models.Recipes, err error)
	DeleteRecipe(id int) (err error)
}
