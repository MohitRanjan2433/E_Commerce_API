package Controllers

import (

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mohit.com/ecom-api/models"
)

// Helper function to parse ObjectID from string
func parseObjectID(idStr string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(idStr)
}

// CreateInventoryController handles the creation of inventory records
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
	productID, err := parseObjectID(request.ProductID)
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

// GetAllInventoryController retrieves all inventory records
func GetAllInventoryController(c *fiber.Ctx) error {
	// Fetch inventory records
	inventories, err := models.GetAllInventory()
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

// UpdateStockController updates the stock of a specific product
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
	productID, err := parseObjectID(request.ProductID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product ID",
		})
	}

	// Update stock
	err = models.UpdateStock(productID, request.Stock)
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

