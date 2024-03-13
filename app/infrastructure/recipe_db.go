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
