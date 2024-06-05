package main

import (
	"go-boilerplate/cmd/internal"
	"go-boilerplate/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {
	utils.InitEnv()

	service := internal.InitServices()

	server := fiber.New()
	httpService := internal.InitHttpService(service)

	app := server.Group("/api/v1")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})

	httpService.Todo.SetRoute(app.Group("/todos"))

	server.Listen(":3000")
}
