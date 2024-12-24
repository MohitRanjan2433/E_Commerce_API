package inventory

import (
	"github.com/gofiber/fiber/v2"
	parseHelper "mohit.com/ecom-api/helper"
	"mohit.com/ecom-api/service"
)

func UpdateStockController(c *fiber.Ctx) error {
	var request struct {
		ProductID string `json:"product_id"`
		Stock     int    `json:"stock"`
	}

	// Parse request body
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error parsing request body",
		})
	}

	// Validate stock
	if request.Stock < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Stock must be greater than zero",
		})
	}

	// Validate product ID
	productID, err := parseHelper.ParseObjectID(request.ProductID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product ID",
		})
	}

	// Update stock
	err = service.UpdateStock(productID, request.Stock)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update stock: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Stock updated successfully",
		"data":    request.Stock,
	})
}
