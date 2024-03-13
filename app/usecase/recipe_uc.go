package usecase

import (
	"backend-test/app/port/adapter"
	"backend-test/models"
	"fmt"

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

func (r *recipeUC) UpdateRecipe(id int, input *models.Recipes) (output *models.Recipes, err error) {
	existing, err := r.repo.GetRecipeByID(id)
	if err != nil {
		log.WithFields(log.Fields{"request": input, "usecase": "UpdateRecipe", "error": err.Error()}).Info("Error while save data")
		return
	}

	if existing == nil {
		err = fmt.Errorf("data not found id: %d", id)
		log.WithFields(log.Fields{"request": input, "usecase": "UpdateRecipe", "error": err.Error()}).Info("Error while save data")
		return
	}

	output, err = r.repo.UpdateRecipe(id, input)
	if err != nil {
		log.WithFields(log.Fields{"request": input, "usecase": "UpdateRecipe", "error": err.Error()}).Info("Error while save data")
		return
	}

	return
}

func (r *recipeUC) DeleteRecipe(id int) (err error) {
	err = r.repo.DeleteRecipe(id)
	if err != nil {
		log.WithFields(log.Fields{"request": id, "usecase": "DeleteRecipe", "error": err.Error()}).Info(fmt.Sprintf("Error while delete data"))
		return
	}
	return
}

func (r *recipeUC) GetListRecipe(request *models.RequestListRecipe) (output *models.ResponseListRecipe, err error) {
	recipes, count, err := r.repo.ListRecipe(request)
	if err != nil {
		return
	}

	paging := models.Paging{
		Page:    request.Page,
		PerPage: request.PerPage,
		Counter: count,
	}

	output = &models.ResponseListRecipe{
		Data:   recipes,
		Paging: paging,
	}

	return
}
