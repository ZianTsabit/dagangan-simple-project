package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ZianTsabit/dagangan-simple-project/handlers"
)

func setupRoutes(app *fiber.App) {
	
	api.Get("/", handlers.Home)

}