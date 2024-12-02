package Controllers

import (
	"github.com/gofiber/fiber/v2"
	"mohit.com/ecom-api/models"
)

func CreateBrand (c *fiber.Ctx) error {
	var request struct{
		Name string `json:"name"`
		Country string `json:"country"`
	}

	if err := c.BodyParser(&request); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":"Invalid Body Request",
		})
	}

	if request.Name == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Brand name is required",
        })
    }

	exists, err := models.CheckIfBrandExists(request.Name)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error checking brand existence",
		})
	}

	if exists {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Brand already exists",
		})
	}

	BrandID, err := models.CreateBrand(request.Name, request.Country)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error creating brand",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Brand created successfully",
		"brand_id": BrandID.Hex(),
	})

}