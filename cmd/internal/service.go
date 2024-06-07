package internal

import (
	repoTodo "go-boilerplate/internal/repository/todo"
	"go-boilerplate/internal/service"
	servTodo "go-boilerplate/internal/service/todo"
)

type Service struct {
	TodoService service.Todo
}

func InitServices() *Service {
	db := InitDB()

	// repository
	todoRepo := repoTodo.NewTodoRepo(db)

	// service
	todoService := servTodo.NewTodoService(servTodo.TodoServiceConfig{
		TodoRepo: todoRepo,
	})

	serv := &Service{
		TodoService: todoService,
	}

	return serv
}
