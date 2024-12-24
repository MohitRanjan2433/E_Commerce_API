package product

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mohit.com/ecom-api/service"
)

func CreateProduct(c *fiber.Ctx) error {
	var request struct {
		Name        string             `json:"name"`
		Description string             `json:"description"`
		Price       float64            `json:"price"`
		Rating      float64            `json:"rating"`
		CategoryID  primitive.ObjectID `json:"category_id"`
		BrandID     primitive.ObjectID `json:"brand_id"`
		Stock       int                `json:"stock"`
	}

	// Parse the body of the request into the struct
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Could not parse the request body",
		})
	}

	// Validate the incoming product data
	if request.Name == "" || request.Price <= 0 || request.Stock < 0 || request.CategoryID.IsZero() || request.BrandID.IsZero() {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product data",
		})
	}

	// Check if product already exists by name and brand
	exists, err := service.CheckIfProductExists(request.Name, request.BrandID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to check if the product exists",
		})
	}

	if exists {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Product with the same name and brand already exists",
		})
	}

	// Call the model function to create the product
	err = service.CreateProduct(request.Name, request.Description, request.Price, request.Rating, request.CategoryID, request.BrandID, request.Stock)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not create the product",
		})
	}

	// Return success response
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Product created successfully",
		"product": request,
	})
}