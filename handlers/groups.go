package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func HandleGroups(c *fiber.Ctx) error {
	return c.Render("groups/index", fiber.Map{})
}
