package Controllers

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mohit.com/ecom-api/models"
)


func GetAllProducts(c *fiber.Ctx) error {
	products, err := models.GetAllProducts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not retrieve products"})
	}
	return c.Status(fiber.StatusOK).JSON(products)
}


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
	exists, err := models.CheckIfProductExists(request.Name, request.BrandID)
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
	err = models.CreateProduct(request.Name, request.Description, request.Price, request.Rating, request.CategoryID, request.BrandID, request.Stock)
	if err != nil {
		// Handle the conflict error or any other error
		if _, ok := err.(*models.ProductConflictError); ok {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

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



func GetProductByPID(c *fiber.Ctx) error {
	pid := c.Params("pid")
	if pid == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "PID is required"})
	}

	product, err := models.GetProductByPID(pid)
	if err != nil {
		if _, ok := err.(*models.ProductNotFoundError); ok {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not retrieve product"})
	}

	return c.Status(fiber.StatusOK).JSON(product)
}

// UpdateProductByPID handles the HTTP request for updating a product by PID
func UpdateProductByPID(c *fiber.Ctx) error {
	pid := c.Params("pid")
	var updatedProduct models.Product
	if err := c.BodyParser(&updatedProduct); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Could not parse the request body"})
	}

	err := models.UpdateProductByPID(pid, updatedProduct)
	if err != nil {
		if _, ok := err.(*models.ProductNotFoundError); ok {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not update product"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Product updated successfully"})
}

// DeleteProductByPID handles the HTTP request for deleting a product by PID
func DeleteProductByPID(c *fiber.Ctx) error {
	pid := c.Params("pid")
	err := models.DeleteProductByPID(pid)
	if err != nil {
		if _, ok := err.(*models.ProductNotFoundError); ok {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not delete product"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Product deleted successfully"})
}
