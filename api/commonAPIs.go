package api

import (
	"mock-server-fiber/service"

	"github.com/gofiber/fiber/v2"
)

func SetCommonAPIs(app *fiber.App) {
	rounter := app.Group("/")

	rounter.Get("/health", service.CheckHealth)
	rounter.Put("/api/v1/sleeptime/:seconds", service.SetSleepTimeInSeconds)
}
