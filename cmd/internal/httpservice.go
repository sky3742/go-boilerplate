package internal

import "go-boilerplate/internal/httpservice/todo"

type HttpService struct {
	Todo *todo.TodoHandler
}

func InitHttpService(service *Service) HttpService {
	return HttpService{
		Todo: todo.NewHandler(todo.ConfigHandler{
			TodoService: service.TodoService,
		}),
	}
}
