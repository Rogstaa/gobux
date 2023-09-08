package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func HandleHelpers(c *fiber.Ctx) error {
	return c.Render("helpers/index", fiber.Map{})
}
