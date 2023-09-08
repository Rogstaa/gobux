package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func HandleDashboard(c *fiber.Ctx) error {
	return c.Render("dashboard/index", fiber.Map{})
}
