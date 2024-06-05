package todo

import (
	"go-boilerplate/internal/repository"
	todoRepo "go-boilerplate/internal/repository/todo"
	"go-boilerplate/internal/service"

	"gorm.io/gorm"
)

type todoService struct {
	TodoRepo repository.TodoProvider
}

type TodoServiceConfig struct {
	TodoRepo repository.TodoProvider
}

func NewTodoService(cfg TodoServiceConfig) service.Todo {
	return &todoService{
		TodoRepo: cfg.TodoRepo,
	}
}

func InitTodoService(db *gorm.DB) service.Todo {
	return NewTodoService(TodoServiceConfig{
		TodoRepo: todoRepo.NewTodoRepo(db),
	})
}
