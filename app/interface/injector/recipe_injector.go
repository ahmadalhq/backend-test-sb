package injector

import (
	"backend-test/app/infrastructure"
	"backend-test/app/interface/api"
	"backend-test/app/usecase"

	"gorm.io/gorm"
)

func RecipeInjector(db *gorm.DB) *api.RecipeApi {
	dbConnection := infrastructure.NewConnectionDB(db)
	recipeUsecase := usecase.RecipeUsecase(dbConnection)
	recipeApi := api.NewRecipeApi(recipeUsecase)
	return recipeApi
}
