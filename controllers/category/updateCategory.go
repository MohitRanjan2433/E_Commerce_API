package category

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"mohit.com/ecom-api/service"
)


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

    err := service.UpdateCategory(categoryID, request.Name)
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