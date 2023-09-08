package main

import (
	// My packages
	//"github.com/Rogstaa/gobux/data"
	"github.com/Rogstaa/gobux/handlers"
	//	"github.com/Rogstaa/gobux/data"
	//"github.com/Rogstaa/gobux/database"

	//fiber
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
	//"log"
)

func main() {
	println("Starting server...")
	// database.ConnectDb()
	//  err := data.InsertProduct("ABC123", 100)
	//  if err != nil {
	//    log.Printf("Error inserting product: %v", err)
	//  } else {
	//    log.Println("Product inserted successfully")
	//  }
	//
	engine := django.New("./views", ".html")
	engine.Reload(true)

	app := fiber.New(fiber.Config{
	ErrorHandler:          handlers.ErrorHandler,
	DisableStartupMessage: true,
	PassLocalsToViews:     true,
	Views:                 engine,

	})

//	data.GetProduct()
	initRoutes(app)

	

	app.Listen(":3000")

}

func initRoutes(app *fiber.App) {


	app.Get("/", handlers.HandleHome)
	app.Get("/admin", handlers.HandleHome)
	app.Get("/groups", handlers.HandleGroups)
	app.Get("/helpers", handlers.HandleHelpers)
	app.Get("/dashboard", handlers.HandleDashboard)
	app.Get("/test", handlers.HandleDashboard)
	app.Get("/admin", handlers.HandleAdmin)
	app.Get("/organization", handlers.HandleOrganization)
	app.Get("/activities", handlers.HandleActivities)


	app.Use(handlers.NotFoundMiddleware) 
	app.Static("/public", "./public")
}


