package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func HandleActivities(c *fiber.Ctx) error {
	return c.Render("activities/index", fiber.Map{})
}



