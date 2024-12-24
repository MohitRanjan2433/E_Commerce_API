package brand

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mohit.com/ecom-api/models"
)

func UpdateBrandController(c *fiber.Ctx) error {
	brandIDHex := c.Params("id")

	brandID, err := primitive.ObjectIDFromHex(brandIDHex)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Brand ID",
		})
	}

	var request struct{
		Name      string    `json:"name"`
		Country   string    `json:"country"`
	}

	if err := c.BodyParser(&request); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":"Invalid request Body",
		})
	}

	err = models.UpdateBrand(brandID, request.Name, request.Country)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":"Error updating brand",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":"Brand Updated Successfully",
	})
}