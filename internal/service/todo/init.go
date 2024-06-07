package todo

import (
	"go-boilerplate/internal/repository"
	"go-boilerplate/internal/service"
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
