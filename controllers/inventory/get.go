package inventory

import (
	"github.com/gofiber/fiber/v2"
	"mohit.com/ecom-api/service"
)

func GetAllInventoryController(c *fiber.Ctx) error {
	// Fetch inventory records
	inventories, err := service.GetAllInventory()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve inventory: " + err.Error(),
		})
	}

	if len(inventories) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No inventory records found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": inventories,
	})
}