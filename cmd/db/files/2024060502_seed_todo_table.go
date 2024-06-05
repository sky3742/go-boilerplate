package files

import (
	"context"
	"go-boilerplate/internal/model"
	"go-boilerplate/internal/service/todo"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func SeedTodoTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "2022081801_seed_todo_table",

		Migrate: func(db *gorm.DB) error {
			todoService := todo.InitTodoService(db)

			_, err := todoService.CreateTodo(context.Background(), &model.Todo{
				Title:     "Todo 1",
				Completed: false,
			})

			return err
		},
	}
}
