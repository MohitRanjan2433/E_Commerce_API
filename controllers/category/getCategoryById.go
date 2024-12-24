package category

import (
	"github.com/gofiber/fiber/v2"
	"mohit.com/ecom-api/service"
)


func GetCategoryByID(c *fiber.Ctx) error {

    categoryID := c.Params("id")

    category, err := service.GetCategoryByID(categoryID)

    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to get category by ID",
        })
    }

    if category == nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Category not found",
        })
    }

    return c.Status(fiber.StatusOK).JSON(category)
}