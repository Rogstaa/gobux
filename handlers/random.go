package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func HandleRandom(c *fiber. Ctx) error {
	return c.Render("random/index", fiber.Map{})
}
