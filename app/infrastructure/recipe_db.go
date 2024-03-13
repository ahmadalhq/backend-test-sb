package infrastructure

import (
	"backend-test/app/port/adapter"
	"backend-test/models"
	"fmt"

	"gorm.io/gorm"
)

type connectionDB struct {
	db *gorm.DB
}

func NewConnectionDB(db *gorm.DB) adapter.RecipeOutputAdapter {
	return &connectionDB{db: db}
}

func (c *connectionDB) InsertRecipe(input *models.Recipes) (output *models.Recipes, err error) {
	tx := c.db.Begin()
	defer func() {
		if r := recover(); r != nil || err != nil {
			if err == nil {
				err = fmt.Errorf("error insert DB, recover from panic: %v", r)
			}
			tx.Rollback()
		}
	}()

	err = c.db.Create(&input).Error
	if err != nil {
		return
	}

	err = tx.Commit().Error
	if err != nil {
		return
	}

	return input, err
}

func (c *connectionDB) UpdateRecipe(id int, input *models.Recipes) (output *models.Recipes, err error) {
	tx := c.db.Begin()
	defer func() {
		if r := recover(); r != nil || err != nil {
			if err == nil {
				err = fmt.Errorf("error insert DB, recover from panic: %v", r)
			}
			tx.Rollback()
		}
	}()

	err = c.db.Where("id = ?", id).Updates(&input).Error
	if err != nil {
		return
	}

	err = tx.Commit().Error
	if err != nil {
		return
	}
	return input, err
}

func (c *connectionDB) DeleteRecipe(id int) (err error) {
	tx := c.db.Begin()
	defer func() {
		if r := recover(); r != nil || err != nil {
			if err == nil {
				err = fmt.Errorf("error insert DB, recover from panic: %v", r)
			}
			tx.Rollback()
		}
	}()

	var existing *models.Recipes
	err = c.db.Where("id = ?", id).Find(&existing).Error
	if err != nil {
		return
	}

	if existing == nil {
		err = fmt.Errorf("data not exist")
		return
	}

	err = c.db.Where("id = ?", id).Delete(&existing).Error
	if err != nil {
		return
	}

	return tx.Commit().Error
}

func (c *connectionDB) ListRecipe(request *models.RequestListRecipe) (output []*models.Recipes, count int64, err error) {
	query := c.db.Model(&models.Recipes{})
	query = query.Where("recipes.name ILIKE ?", "%"+request.Params+"%")

	if request.OrderBy != "" {
		if request.SortType != "" {
			query = query.Order(request.OrderBy + " " + request.SortType)
		} else {
			query = query.Order(request.OrderBy)
		}
	}

	err = query.Count(&count).Error
	if err != nil {
		return
	}

	if request.Page > 0 && request.PerPage > 0 {
		offset := (request.Page - 1) * request.PerPage
		query = query.Limit(request.PerPage).Offset(offset)
	}

	err = query.Find(&output).Error
	if err != nil {
		return
	}

	return
}

func (c *connectionDB) GetRecipeByID(id int) (out *models.Recipes, err error) {
	err = c.db.Where("id = ?", id).Find(&out).Error
	return
}
