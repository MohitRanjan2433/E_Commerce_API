package brand

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mohit.com/ecom-api/models"
)



func DeleteBrandController(c *fiber.Ctx) error {
	brandIDHex := c.Params("id")

	brand, err := primitive.ObjectIDFromHex(brandIDHex)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":"Invalid brand id",
		})
	}

	err = models.DeleteBrand(brand)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error deleting brand",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Brand Deleted Successfully",
	})
}