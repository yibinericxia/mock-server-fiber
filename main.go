package main

import (
	"mock-server-fiber/api"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	api.SetCommonAPIs(app)
	api.SetupUserAPIs(app)

	app.Listen(":3000")
}
