package main

import (
	"go-boilerplate/cmd/internal"
	"go-boilerplate/internal/middleware"
	"go-boilerplate/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	utils.InitEnv()

	server := fiber.New()
	server.Use(logger.New())
	server.Use(healthcheck.New())
	server.Use(recover.New())

	app := server.Group("/api/v1")

	service := internal.InitServices()
	httpService := internal.InitHttpService(service)

	httpService.Todo.SetPublicRoute(app.Group("/todos"))
	httpService.Todo.SetPrivateRoute(app.Group("/admin/todos"), middleware.SessionMiddleware)

	server.Listen(":3000")
}
