package brand

import (
	"github.com/gofiber/fiber/v2"
	"mohit.com/ecom-api/service"
)


func GetAllBrandController(c *fiber.Ctx) error {
	brands, err := service.GetAllBrand()
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