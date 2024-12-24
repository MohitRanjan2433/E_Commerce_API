package cart

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"mohit.com/ecom-api/service"
)

func RemoveItemFromCart(c *fiber.Ctx) error {
	// Retrieve the product_id from the URL parameter
	productID := c.Params("product_id")

	// Ensure userID exists in the context
	userID, ok := c.Locals("userID").(string)
	if !ok || userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "UserID is missing in the context. Please authenticate.",
		})
	}

	// Log the productID and userID for debugging
	fmt.Println("Removing item with product_id:", productID, "for user:", userID)

	// Remove the item from the cart
	err := service.RemoveItemFromCart(userID, productID)
	if err != nil {
		// Log error details for debugging
		fmt.Println("Error removing item from cart:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not remove item from cart", "details": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Item removed from cart"})
}