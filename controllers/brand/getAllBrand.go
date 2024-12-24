package brand

import (
	"github.com/gofiber/fiber/v2"
	"mohit.com/ecom-api/models"
)


func GetAllBrandController(c *fiber.Ctx) error {
	brands, err := models.GetAllBrand()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error fetching brands",
		})
	}

	// Return the list of brands as a JSON response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"brands": brands,
	})

}