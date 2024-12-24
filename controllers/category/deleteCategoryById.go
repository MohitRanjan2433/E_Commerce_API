package category

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"mohit.com/ecom-api/models" 
)

func DeleteCategoryByID(c *fiber.Ctx) error {
	id := c.Params("id") 
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Category ID is required"})
	}

	err := models.DeleteCategoryByID(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Category not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not delete category"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Category deleted successfully"})
}