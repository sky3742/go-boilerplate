package repository

import (
	"context"
	"go-boilerplate/internal/model"
)

type TodoProvider interface {
	GetTodos(ctx context.Context) ([]model.Todo, error)
	GetTodo(ctx context.Context, id string) (*model.Todo, error)
	CreateTodo(ctx context.Context, todo *model.Todo) (*model.Todo, error)
	UpdateTodo(ctx context.Context, todo *model.Todo) (*model.Todo, error)
	DeleteTodo(ctx context.Context, id string) error
}
