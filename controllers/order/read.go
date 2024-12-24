package order

import (
	"github.com/gofiber/fiber/v2"
	"mohit.com/ecom-api/models"
)


func GetAllOrdersController(c *fiber.Ctx) error {

	userID, ok := c.Locals("userID").(string)
	if !ok || userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "UserID is missing in the context. Please authenticate.",
		})
	}

	orders, err := models.GetAllOrdersById(userID)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve orders: " + err.Error(),
		})
	}

	if orders == nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No orders found for the user",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"orders": orders,
	})
}