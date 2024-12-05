package api

import (
	"mock-server-fiber/service"

	"github.com/gofiber/fiber/v2"
)

func SetupUserAPIs(app *fiber.App) {
	base := app.Group("/api/v1")
	router := base.Group("/users")

	router.Get("/", service.FindUsers)
	router.Get("/:id", service.FindUserByID)
	router.Post("/", service.CreateUser)
	router.Put("/:id", service.UpdateUser)
	router.Delete("/:id", service.DeleteUserByID)
}