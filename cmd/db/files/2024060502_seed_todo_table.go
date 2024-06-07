package files

import (
	"go-boilerplate/internal/model"
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func SeedTodoTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "2022081801_seed_todo_table",

		Migrate: func(db *gorm.DB) error {
			result := db.Create(&[]model.Todo{
				{
					Title:     "Todo 1",
					Completed: false,
				},
				{
					Title:     "Todo 2",
					Completed: true,
				},
			})

			log.Println("seeding todos table")
			return result.Error
		},

		Rollback: func(d *gorm.DB) error {
			log.Println("skipping rollback seed todos table")
			return nil
		},
	}
}
