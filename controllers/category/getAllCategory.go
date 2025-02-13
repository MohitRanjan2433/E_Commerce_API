package category

import (
	"github.com/gofiber/fiber/v2"
	"mohit.com/ecom-api/service"
)


func GetAllCategory(c *fiber.Ctx) error {
	categories, err := service.GetAllCategory()
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get all categories",
		})
	}

	return c.Status(fiber.StatusOK).JSON(categories)
}