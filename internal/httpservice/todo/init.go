package todo

import (
	"go-boilerplate/internal/service"

	"github.com/gofiber/fiber/v2"
)

type TodoHandler struct {
	TodoService service.Todo
}

type ConfigHandler struct {
	TodoService service.Todo
}

func NewHandler(cfg ConfigHandler) *TodoHandler {
	return &TodoHandler{
		TodoService: cfg.TodoService,
	}
}

func (h *TodoHandler) SetRoute(app fiber.Router) {
	app.Get("/", h.GetTodos)
	app.Get("/:id", h.GetTodo)
	app.Post("/", h.CreateTodo)
	app.Put("/:id", h.UpdateTodo)
	app.Delete("/:id", h.DeleteTodo)
}
