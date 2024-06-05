package todo

import (
	"go-boilerplate/internal/repository"

	"gorm.io/gorm"
)

type todoRepo struct {
	db *gorm.DB
}

func NewTodoRepo(db *gorm.DB) repository.TodoProvider {
	return &todoRepo{
		db: db,
	}
}
