package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ZianTsabit/dagangan-simple-project/database"
	"github.com/ZianTsabit/dagangan-simple-project/models"
	"strconv"
)

// Get Products using filter and pagination
func GetProducts(c *fiber.Ctx) error {
	var products []models.Product
	sql := "SELECT * FROM products"

	// Filter by product name
	if c.Query("name") != "" {
		sql += " WHERE name LIKE '%" + c.Query("name") + "%'"
	}

	// Filter by product price
	if c.Query("price") != "" {
		if c.Query("name") != "" {
			sql += " AND price = " + c.Query("price")
		} else {
			sql += " WHERE price = " + c.Query("price")
		}
	}

	sql += " ORDER BY id"

	// Page size
	page_size := 5
	if c.Query("page_size") != "" {
		sql += " LIMIT " + c.Query("page_size")
		page_size, _ = strconv.Atoi(c.Query("page_size"))
	} else {
		sql += " LIMIT 5"
	}

	// Pagination
	if c.Query("page_num") != "" {
		page_num, _ := strconv.Atoi(c.Query("page_num"))
		page := (page_num - 1) * page_size
		sql += " OFFSET " + strconv.Itoa(page)
	} else {
		sql += " OFFSET 0"
	}

	// Execute query
	database.Db.Db.Raw(sql).Scan(&products)
	
	//just return the products id, name, price, quantity
	data := make([]map[string]interface{}, len(products))
	for i, product := range products {
		data[i] = map[string]interface{}{
			"id": product.Id,
			"name": product.Name,
			"price": product.Price,
			"quantity": product.Quantity,
		}
	}
	
	return c.JSON(
		fiber.Map{
			"status": "success",
			"message": "Successfully retrieved products",
			"data": data,
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

	sql := "DELETE FROM products WHERE id = " + Id
	database.Db.Db.Exec(sql)

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
	newPrice := strconv.FormatFloat(product.Price, 'f', 2, 64)
	sql := "UPDATE products SET name = '" + product.Name + "', price = " + newPrice + ", quantity = " + strconv.Itoa(product.Quantity) + " WHERE id = " + Id
	database.Db.Db.Exec(sql)
	
	return c.JSON(
		fiber.Map{
			"status": "success",
			"message": "Successfully updated product",
			"data": product,
		},
	)
}