package cart

import (
	"github.com/gofiber/fiber/v2"
	"mohit.com/ecom-api/service"
)

func AddItemToCart(c *fiber.Ctx) error {
	var request struct {
		ProductID   string  `json:"product_id"`
		ProductName string  `json:"product_name"` // Add Product Name here
		Quantity    int     `json:"quantity"`
		Price       float64 `json:"price"`
	}

	// Parse the request body
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request", "details": err.Error()})
	}

	// Ensure userID exists in the context
	userID, ok := c.Locals("userID").(string)
	if !ok || userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "UserID is missing in the context. Please authenticate.",
		})
	}

	// Add the item to the cart
	err := service.AddItemToCart(userID, request.ProductID, request.ProductName, request.Quantity, request.Price)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not add item to cart", "details": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Item added to cart"})
}