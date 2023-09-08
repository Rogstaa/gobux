// handlers.go
package handlers

import (
	"log"
	//"text/template"
	"github.com/Rogstaa/gobux/data"  // Adjust the import path accordingly
	"github.com/gofiber/fiber/v2"
	//"bytes"
)
/*
func HandleAdmin(c *fiber.Ctx) error {
	product := data.GetProduct()
	log.Printf("Product: %+v\n", product) // Debug line

	return c.Render("admin/index", fiber.Map{
		"Code":  product.Code,
		"Price": product.Price,
	})
}
*/

func HandleAdmin(c *fiber.Ctx) error {
	product, err := data.GetProduct()
	if err != nil {
		log.Printf("Error fetching product: %v", err)
		// Handle error, maybe return a status code or message to the client
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}
	log.Printf("Product: %+v\n", product.Code) // Debug line

	return c.Render("admin/index", fiber.Map{
		//"Product": fiber.Map{"Code": "TestCode", "Price": 999},
		"Code":  product.Code,
	"Price": product.Price,
	})
}
