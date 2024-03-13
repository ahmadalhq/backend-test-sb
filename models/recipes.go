package models

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Recipes struct {
	ID          int64          `json:"Id"`
	Name        string         `json:"name"`
	Description string         `json:"desc"`
	Ingredients pq.StringArray `json:"ingredients"`
	Seasonings  pq.StringArray `json:"seasonings"`
	HowTo       pq.StringArray `json:"howTo"`
	CreatedAt   *time.Time     `json:"createdAt,omitempty"`
	UpdatedAt   *time.Time     `json:"updatedAt,omitempty"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt,omitempty" sql:"index"`
}
