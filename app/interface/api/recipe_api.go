package api

import (
	"backend-test/app/port/adapter"
	"backend-test/models"
	"backend-test/utils"
	"fmt"
	"net/http"
	"strconv"

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
		log.WithFields(log.Fields{"error": err.Error()}).Error(fmt.Sprintf("Error while binding request %s", "Recipe"))
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

func (r *RecipeApi) UpdateRecipe(ctx echo.Context) (err error) {
	data := new(models.Recipes)
	id, _ := strconv.Atoi(ctx.Param("id"))

	err = ctx.Bind(data)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error()}).Info(fmt.Sprintf("Error while binding request %s", "Recipe"))
		return utils.SendResponse(ctx, utils.ResponseMessage{
			Code:    http.StatusConflict,
			Message: err.Error(),
		})
	}

	output, err := r.repo.UpdateRecipe(id, data)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error()}).Info(fmt.Sprintf("Error while binding request %s", "Recipe"))
		return utils.SendResponse(ctx, utils.ResponseMessage{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return utils.SendResponse(ctx, utils.ResponseMessage{
		Code:    http.StatusOK,
		Message: "success",
		Data:    output,
	})
}

func (r *RecipeApi) GetListRecipe(ctx echo.Context) (err error) {
	request := new(models.RequestListRecipe)
	err = ctx.Bind(request)
	resp := new(models.ResponseListRecipe)

	if err != nil {
		log.WithFields(log.Fields{"error": err.Error()}).Error(fmt.Sprintf("Error while binding request %s", "Recipe"))
		return utils.SendResponse(ctx, utils.ResponseMessage{
			Code:    http.StatusConflict,
			Message: err.Error(),
		})
	}

	if request.Page == 0 {
		request.Page = 1
	}
	if request.PerPage == 0 {
		request.PerPage = 10
	}

	resp, err = r.repo.GetListRecipe(request)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error()}).Error(fmt.Sprintf("Error while binding request %s", "Recipe"))
		return utils.SendResponse(ctx, utils.ResponseMessage{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return utils.SendResponse(ctx, utils.ResponseMessage{
		Code:    http.StatusOK,
		Message: "success",
		Data:    resp,
	})
}

func (r *RecipeApi) DeleteRecipe(ctx echo.Context) (err error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err = r.repo.DeleteRecipe(id)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error()}).Info(fmt.Sprintf("Error while doing request %s", "Recipe"))
		return utils.SendResponse(ctx, utils.ResponseMessage{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return utils.SendResponse(ctx, utils.ResponseMessage{
		Code:    http.StatusOK,
		Message: "success",
		Data:    id,
	})
}
