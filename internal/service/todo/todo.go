package todo

import (
	"context"
	"go-boilerplate/internal/model"
)

func (s *todoService) GetTodos(ctx context.Context) ([]model.Todo, error) {
	todos, err := s.TodoRepo.GetTodos(ctx)
	return todos, err
}

func (s *todoService) GetTodo(ctx context.Context, id string) (model.Todo, error) {
	todo, err := s.TodoRepo.GetTodo(ctx, id)
	return *todo, err
}

func (s *todoService) CreateTodo(ctx context.Context, todo *model.Todo) (model.Todo, error) {
	todo, err := s.TodoRepo.CreateTodo(ctx, todo)
	return *todo, err
}

func (s *todoService) UpdateTodo(ctx context.Context, todo *model.Todo) (model.Todo, error) {
	todo, err := s.TodoRepo.UpdateTodo(ctx, todo)
	return *todo, err
}

func (s *todoService) DeleteTodo(ctx context.Context, id string) error {
	err := s.TodoRepo.DeleteTodo(ctx, id)
	return err
}
