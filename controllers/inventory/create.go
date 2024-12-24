package inventory

import (

	"github.com/gofiber/fiber/v2"
	"mohit.com/ecom-api/models"
	parseHelper"mohit.com/ecom-api/helper"
	
)

func CreateInventoryController(c *fiber.Ctx) error {
	var request struct {
		ProductID string `json:"product_id"`
		Stock     int    `json:"stock"`
		Warehouse string `json:"warehouse"`
	}

	// Parse request body
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error parsing inventory request body",
		})
	}

	// Validate product ID
	productID, err := parseHelper.ParseObjectID(request.ProductID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product ID",
		})
	}

	// Create inventory
	err = models.CreateInventory(productID, request.Warehouse, request.Stock)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error creating inventory: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Inventory created successfully",
	})
}