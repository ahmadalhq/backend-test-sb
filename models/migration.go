package models

import (
	cm "backend-test/common"
)

func Migrate() {
	cm.GetInstancePostgresDb().AutoMigrate(
		&Recipes{}, //create table recipes
	)
}
