package Controllers

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mohit.com/ecom-api/models"
)



func CreateInventoryController(c *fiber.Ctx) error {
	var request struct{
		ProductID    string    `json:"product_id"`
		Stock        int       `json:"stock"`
		Warehouse    string    `json:"warehouse"`
	}

	if err := c.BodyParser(&request); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors":"Error in parsing inventory body request",
		})
	}

	// Ensure userID exists in the context
	userID, ok := c.Locals("userID").(string)
	if !ok || userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "UserID is missing in the context. Please authenticate.",
		})
	}

	err := models.CreateInventory(userID, request.ProductID,request.Warehouse, request.Stock)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errors":"Error in calling invent model",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":"Succesfully inventory created",
	})

}

func GetAllInventoryCollection(c *fiber.Ctx) error {
	
	// Ensure userID exists in the context
	userID, ok := c.Locals("userID").(string)
	if !ok || userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "UserID is missing in the context. Please authenticate.",
		})
	}

	inventory, err := models.GetAllInventory(userID)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve orders" + err.Error(),
		})
	}

	if inventory == nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No data found in inventory",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": inventory,
	})
}

func UpdateStockController(c *fiber.Ctx) error {
	var request struct{
		ID         primitive.ObjectID   `json:"_id"`
		ProductID  string  `json:"product_id"`
		Stock     int     `json:"stock"`
	}

	if request.Stock < 1 {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"error": "Stock cannot be 0 or less",
		})
	}

	// Ensure userID exists in the context
	userID, ok := c.Locals("userID").(string)
	if !ok || userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "UserID is missing in the context. Please authenticate.",
		})
	}

	if err := c.BodyParser(&request); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error parsing the body req",
		})
	}

	if request.Stock == 0{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Stock in body is empty",
		})
	}

	err := models.UpdateStock(userID, request.ProductID, request.ID, request.Stock)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update order status" + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Stock updated successfully",
		"data": request.Stock,
	})

}

