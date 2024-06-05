package todo

import (
	"go-boilerplate/internal/model"

	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
)

func (h *TodoHandler) GetTodos(c *fiber.Ctx) error {
	todos, err := h.TodoService.GetTodos(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse{
			StatusCode: fiber.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.JsonResponse{
		StatusCode: fiber.StatusOK,
		Message:    "success get todos",
		Data:       todos,
	})
}

func (h *TodoHandler) GetTodo(c *fiber.Ctx) error {
	todo, err := h.TodoService.GetTodo(c.Context(), c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse{
			StatusCode: fiber.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.JsonResponse{
		StatusCode: fiber.StatusOK,
		Message:    "success get todo",
		Data:       &todo,
	})
}

func (h *TodoHandler) CreateTodo(c *fiber.Ctx) error {
	var todo model.Todo

	err := c.BodyParser(&todo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			StatusCode: fiber.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	todo, err = h.TodoService.CreateTodo(c.Context(), &todo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse{
			StatusCode: fiber.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.JsonResponse{
		StatusCode: fiber.StatusOK,
		Message:    "success create todo",
		Data:       todo,
	})
}

func (h *TodoHandler) UpdateTodo(c *fiber.Ctx) error {
	var todo model.Todo

	err := c.BodyParser(&todo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			StatusCode: fiber.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	todo.ID = uuid.FromStringOrNil(c.Params("id"))

	todo, err = h.TodoService.UpdateTodo(c.Context(), &todo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse{
			StatusCode: fiber.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.JsonResponse{
		StatusCode: fiber.StatusOK,
		Message:    "success update todo",
		Data:       todo,
	})
}

func (h *TodoHandler) DeleteTodo(c *fiber.Ctx) error {
	err := h.TodoService.DeleteTodo(c.Context(), c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse{
			StatusCode: fiber.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.JsonResponse{
		StatusCode: fiber.StatusOK,
		Message:    "success delete todo",
	})
}
