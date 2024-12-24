package category

import (
	"github.com/gofiber/fiber/v2"
	"mohit.com/ecom-api/service"
)


func CreateCategory(c *fiber.Ctx) error {
    var request struct {
        Name string `json:"name"`
    }

    // Parse the body of the request
    if err := c.BodyParser(&request); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    if request.Name == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Category name is required",
        })
    }

    // Check if category already exists
    exists, err := service.CheckIfCategoryExists(request.Name)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to check if category exists",
        })
    }

    if exists {
        return c.Status(fiber.StatusConflict).JSON(fiber.Map{
            "error": "Category already exists",
        })
    }

    // Create the category
    categoryID, err := service.CreateCategory(request.Name)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to create category",
        })
    }

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "message":     "Category created successfully",
        "category_id": categoryID.Hex(),
    })
}