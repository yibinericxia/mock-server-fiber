package test

import (
	"mock-server-fiber/service"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestUser(t *testing.T) {
	// ctx := context.Background()
	err := service.FindUsers(&fiber.Ctx{})
	if err != nil {
		t.Error(err)
	}
}
