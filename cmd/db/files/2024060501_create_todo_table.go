package files

import (
	"go-boilerplate/internal/model"
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func CreateTodoTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "2022081801_create_todo_table",

		Migrate: func(db *gorm.DB) error {
			type Todo struct {
				model.ModelBase
				Title     string `gorm:""`
				Completed bool   `gorm:""`
			}

			log.Println("migrating todos table")
			return db.AutoMigrate(&Todo{})
		},

		Rollback: func(db *gorm.DB) error {
			log.Println("rolling back todos table")
			return db.Migrator().DropTable("todos")
		},
	}
}
