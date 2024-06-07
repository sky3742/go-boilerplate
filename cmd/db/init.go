package db

import (
	"go-boilerplate/cmd/db/files"
	"go-boilerplate/internal/utils"

	"github.com/go-gormigrate/gormigrate/v2"
)

var Migrations *gormigrate.Gormigrate

func init() {
	utils.InitEnv()

	db := utils.InitDB(utils.Config.DatabaseUrl)

	Migrations = gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		files.CreateTodoTable(),
		files.SeedTodoTable(),
	})
}
