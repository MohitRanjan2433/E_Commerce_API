package brand

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mohit.com/ecom-api/models"
)


func GetBrandController(c *fiber.Ctx)error {
	brandIDHex := c.Params("id")
	brandID, err := primitive.ObjectIDFromHex(brandIDHex)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":"Invalid brand id",
		})
	}

	brand, err := models.GetBrandByID(brandID)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":"Error retriving brand",
		})
	}

	if brand == nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":"Brand not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"brand": brand,
	})
}