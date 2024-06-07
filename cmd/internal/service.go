package internal

import (
	repoTodo "go-boilerplate/internal/repository/todo"
	"go-boilerplate/internal/service"
	servTodo "go-boilerplate/internal/service/todo"
	"go-boilerplate/internal/utils"
)

type Service struct {
	TodoService service.Todo
}

func InitServices() *Service {
	db := utils.InitDB(utils.Config.DatabaseUrl)

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
