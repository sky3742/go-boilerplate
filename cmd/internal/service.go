package internal

import (
	"go-boilerplate/internal/service"
	"go-boilerplate/internal/service/todo"
)

type Service struct {
	TodoService service.Todo
}

func InitServices() *Service {
	db := InitDB()

	todoService := todo.InitTodoService(db)

	serv := &Service{
		TodoService: todoService,
	}

	return serv
}
