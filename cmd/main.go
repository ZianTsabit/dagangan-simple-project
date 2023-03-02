package main

import (
    "log"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/ZianTsabit/dagangan-simple-project/database"
)

func main() {
    database.ConnectDb()
    app := fiber.New()
    app.Use(cors.New())

    setupRoutes(app)
    
    log.Fatal(app.Listen(":8000"))
}


