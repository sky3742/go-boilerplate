package db

import (
	"go-boilerplate/cmd/db/files"
	"go-boilerplate/cmd/internal"
	"go-boilerplate/internal/utils"

	"github.com/go-gormigrate/gormigrate/v2"
)

var Migrations *gormigrate.Gormigrate

func init() {
	utils.InitEnv()

	db := internal.InitDB()

	Migrations = gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		files.CreateTodoTable(),
		files.SeedTodoTable(),
	})
}
