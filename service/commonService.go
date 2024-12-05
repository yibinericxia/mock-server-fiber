package service

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

var seconds int = 0

func CheckHealth(ctx *fiber.Ctx) error {
	return ctx.SendString("It is up!")
}

func SetSleepTimeInSeconds(ctx *fiber.Ctx) error {
	n, err := strconv.Atoi(ctx.Params("seconds"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "invalid sleep time in seconds",
			"message": err.Error(),
		})
	}
	seconds = n
	return ctx.SendString("")
}
