// controllers/category_controller.go
package Controllers

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"mohit.com/ecom-api/models" // Correct import path
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
    exists, err := models.CheckIfCategoryExists(request.Name)
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
    categoryID, err := models.CreateCategory(request.Name)
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

func GetAllCategory(c *fiber.Ctx) error {
	categories, err := models.GetAllCategory()
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get all categories",
		})
	}

	return c.Status(fiber.StatusOK).JSON(categories)
}

func GetCategoryByID(c *fiber.Ctx) error {

    categoryID := c.Params("id")

    category, err := models.GetCategoryByID(categoryID)

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

func UpdateCategory(c *fiber.Ctx) error {

    var request struct{
        Name string `json:"name"`
    }

    categoryID := c.Params("id")

    if err := c.BodyParser(&request); err != nil{
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    if request.Name == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Name is required",
        })
    }

    err := models.UpdateCategory(categoryID, request.Name)
    if err != nil{
        if err == mongo.ErrNoDocuments{
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
                "error": "Category not found",
            })
        }
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to update category",
        })
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Category Updated Successfully",
        "data": request.Name,
    })
}

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
