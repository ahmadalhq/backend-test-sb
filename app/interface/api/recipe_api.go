package api

import (
	"backend-test/app/port/adapter"
	"backend-test/models"
	"backend-test/utils"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type RecipeApi struct {
	repo adapter.RecipeInputAdapter
}

func NewRecipeApi(repo adapter.RecipeInputAdapter) *RecipeApi {
	return &RecipeApi{repo: repo}
}

func (r *RecipeApi) InsertRecipe(ctx echo.Context) (err error) {
	data := new(models.Recipes)

	err = ctx.Bind(data)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error()}).Info(fmt.Sprintf("Error while binding request %s", "Recipe"))
		return utils.SendResponse(ctx, utils.ResponseMessage{
			Code:    http.StatusConflict,
			Message: err.Error(),
		})
	}

	out, err := r.repo.InsertRecipe(data)
	if err != nil {
		return utils.SendResponse(ctx, utils.ResponseMessage{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return utils.SendResponse(ctx, utils.ResponseMessage{
		Code:    http.StatusOK,
		Message: "success",
		Data:    out,
	})
}
