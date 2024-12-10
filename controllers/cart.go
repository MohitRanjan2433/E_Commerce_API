package Controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"mohit.com/ecom-api/models"
)

// AddItemToCart handles adding an item to the user's cart
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
	err := models.AddItemToCart(userID, request.ProductID, request.ProductName, request.Quantity, request.Price)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not add item to cart", "details": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Item added to cart"})
}

// GetCart retrieves the user's cart
func GetCart(c *fiber.Ctx) error {
	userID, ok := c.Locals("userID").(string)
	if !ok || userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "UserID is missing in the context. Please authenticate.",
		})
	}

	cart, err := models.GetCart(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not retrieve cart", "details": err.Error()})
	}

	if cart == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Cart is empty"})
	}

	return c.Status(fiber.StatusOK).JSON(cart)
}

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
	err := models.RemoveItemFromCart(userID, productID)
	if err != nil {
		// Log error details for debugging
		fmt.Println("Error removing item from cart:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not remove item from cart", "details": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Item removed from cart"})
}


// ClearCart handles clearing the user's cart
func ClearCart(c *fiber.Ctx) error {
	// Ensure userID exists in the context
	userID, ok := c.Locals("userID").(string)
	if !ok || userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "UserID is missing in the context. Please authenticate.",
		})
	}

	// Clear the cart
	err := models.ClearCart(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not clear cart", "details": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Cart cleared successfully"})
}
