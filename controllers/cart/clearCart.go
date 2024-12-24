package cart

import (
	"github.com/gofiber/fiber/v2"
	"mohit.com/ecom-api/service"
)



func ClearCart(c *fiber.Ctx) error {
	// Ensure userID exists in the context
	userID, ok := c.Locals("userID").(string)
	if !ok || userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "UserID is missing in the context. Please authenticate.",
		})
	}

	// Clear the cart
	err := service.ClearCart(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not clear cart", "details": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Cart cleared successfully"})
}
