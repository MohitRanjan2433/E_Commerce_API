package cart

import (
	"github.com/gofiber/fiber/v2"
	"mohit.com/ecom-api/service"
)

func GetCart(c *fiber.Ctx) error {
	userID, ok := c.Locals("userID").(string)
	if !ok || userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "UserID is missing in the context. Please authenticate.",
		})
	}

	cart, err := service.GetCart(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not retrieve cart", "details": err.Error()})
	}

	if cart == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Cart is empty"})
	}

	return c.Status(fiber.StatusOK).JSON(cart)
}