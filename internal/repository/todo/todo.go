package todo

import (
	"context"
	"go-boilerplate/internal/model"
)

func (r *todoRepo) GetTodos(ctx context.Context) ([]model.Todo, error) {
	var todos []model.Todo
	result := r.db.Find(&todos)
	return todos, result.Error
}

func (r *todoRepo) GetTodo(ctx context.Context, id string) (*model.Todo, error) {
	var todo *model.Todo
	result := r.db.Where("id = ?", id).First(&todo)
	return todo, result.Error
}

func (r *todoRepo) CreateTodo(ctx context.Context, todo *model.Todo) (*model.Todo, error) {
	result := r.db.Create(&todo)
	return todo, result.Error
}

func (r *todoRepo) UpdateTodo(ctx context.Context, todo *model.Todo) (*model.Todo, error) {
	result := r.db.Save(&todo)
	return todo, result.Error
}

func (r *todoRepo) DeleteTodo(ctx context.Context, id string) error {
	result := r.db.Where("id = ?", id).Delete(&model.Todo{})
	return result.Error
}
