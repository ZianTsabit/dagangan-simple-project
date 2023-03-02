package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ZianTsabit/dagangan-simple-project/handlers"
)

func setupRoutes(app *fiber.App) {
	
	app.Get("/product/", handlers.GetProducts)	
	app.Get("/product/:id", handlers.GetProduct)
	app.Post("/product/", handlers.NewProduct)
	app.Delete("/product/:id", handlers.DeleteProduct)
	app.Put("/product/:id", handlers.UpdateProduct)

}