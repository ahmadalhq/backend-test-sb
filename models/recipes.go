package models

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Recipes struct {
	ID             int64          `json:"Id"`
	Name           string         `json:"name"`
	Description    string         `json:"desc"`
	Ingredients    datatypes.JSON `json:"ingredients" gorm:"column:ingredients"`
	HowTo          pq.StringArray `json:"howTo" gorm:"type:text[]"`
	CookingTime    int64          `json:"cookingTime"`
	ServingPortion int64          `json:"servingPortion"`
	CreatedAt      *time.Time     `json:"createdAt,omitempty"`
	UpdatedAt      *time.Time     `json:"updatedAt,omitempty"`
	DeletedAt      gorm.DeletedAt `json:"deletedAt,omitempty" sql:"index"`
}

type RequestListRecipe struct {
	Search string `query:"search"`
	Params string `query:"params"`
	FilterPagination
}

type ResponseListRecipe struct {
	Data   []*Recipes `json:"data"`
	Paging Paging     `json:"paging"`
}
