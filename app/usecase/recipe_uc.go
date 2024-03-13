package usecase

import (
	"backend-test/app/port/adapter"
	"backend-test/models"

	log "github.com/sirupsen/logrus"
)

type recipeUC struct {
	repo adapter.RecipeOutputAdapter
}

func RecipeUsecase(repo adapter.RecipeOutputAdapter) adapter.RecipeInputAdapter {
	return &recipeUC{repo: repo}
}

func (r *recipeUC) InsertRecipe(input *models.Recipes) (output *models.Recipes, err error) {
	output, err = r.repo.InsertRecipe(input)
	if err != nil {
		log.WithFields(log.Fields{"request": input, "usecase": "InsertRecipe", "error": err.Error()}).Info("Error while save data")
		return
	}
	return
}
