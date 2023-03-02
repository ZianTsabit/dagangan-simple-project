package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ZianTsabit/dagangan-simple-project/database"
	"github.com/ZianTsabit/dagangan-simple-project/models"
)

// Get Products using filter and pagination
func GetProducts(c *fiber.Ctx) error {
	var products []models.Product
	sql := "SELECT * FROM products"

	// Filter
	if c.Query("filter") != "" {
		sql += " WHERE name LIKE '%" + c.Query("filter") + "%'" + " OR price LIKE >= + c.Query("filter")"
	}

	// Page size
	if c.Query("page_size") != "" {
		sql += " LIMIT " + c.Query("page_size")
	} else {
		sql += " LIMIT 5"
	}

	// Pagination
	if c.Query("page") != "" {
		sql += " OFFSET " + c.Query("page")
	} else {
		sql += " OFFSET 0"
	}

	// Execute query
	database.Db.D.Raw(sql).Scan(&products)
	return c.JSON(
		fiber.Map{
			"status": "success",
			"message": "Successfully retrieved products",
			"data": products,
		},
	)
}

func GetProduct(c *fiber.Ctx) error {
	Id := c.Params("id")
	var product models.Product
	database.Db.Db.Find(&product, Id)
	return c.JSON(
		fiber.Map{
			"status": "success",
			"message": "Successfully retrieved product",
			"data": product,
		},
	)
}

func NewProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return err
	}
	database.Db.Db.Create(&product)
	return c.JSON(
		fiber.Map{
			"status": "success",
			"message": "Successfully created product",
			"data": product,
		},
	)
}

func DeleteProduct(c *fiber.Ctx) error {
	Id := c.Params("id")
	var product models.Product
	database.Db.Db.First(&product, Id)
	if product.Name == "" {
		return c.Status(500).SendString("No product found with ID")
	}
	database.Db.Db.Delete(&product)
	return c.JSON(
		fiber.Map{
			"status": "success",
			"message": "Successfully deleted product",
		},
	)
}

func UpdateProduct(c *fiber.Ctx) error {
	Id := c.Params("id")
	var product models.Product
	database.Db.Db.First(&product, Id)
	if product.Name == "" {
		return c.Status(500).SendString("No product found with ID")
	}
	if err := c.BodyParser(&product); err != nil {
		return err
	}
	database.Db.Db.Save(&product)
	return c.JSON(
		fiber.Map{
			"status": "success",
			"message": "Successfully updated product",
			"data": product,
		},
	)
}