package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func HandleTest(c *fiber.Ctx) error {
	return c.Render("test/index", fiber.Map{})
}
