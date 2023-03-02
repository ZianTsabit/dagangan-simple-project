package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ZianTsabit/dagangan-simple-project/database"
	"github.com/ZianTsabit/dagangan-simple-project/models"
)

func GetProducts(c *fiber.Ctx) error {
	var products []models.Product
	database.Db.Db.Find(&products)
	return c.JSON(products)
}

func Home(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

