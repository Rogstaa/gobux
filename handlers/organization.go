package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func HandleOrganization(c *fiber.Ctx) error {
	return c.Render("organization/index", fiber.Map{})
}


