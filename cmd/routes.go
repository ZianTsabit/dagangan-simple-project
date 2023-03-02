package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ZianTsabit/dagangan-simple-project/handlers"
)

func setupRoutes(app *fiber.App) {
	
	app.Get("/", handlers.GetProducts)
	app.Get("/:id", handlers.GetProduct)
	app.Post("/", handlers.NewProduct)
	app.Delete("/:id", handlers.DeleteProduct)
	app.Put("/:id", handlers.UpdateProduct)
	
}