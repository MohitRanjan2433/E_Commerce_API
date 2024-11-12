package Controllers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mohit.com/ecom-api/db"
	"mohit.com/ecom-api/models"
)

func GetAllproducts(c *fiber.Ctx) error {
	productCollection :=  db.GetProductCollection()

	//set a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := productCollection.Find(ctx, bson.M{}, options.Find())
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not retieve product",
		})
	}

	defer cursor.Close(ctx)

	var product[] models.Product
	if err := cursor.All(ctx, &product); err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error passing product",
		})
	}

	return c.Status(fiber.StatusOK).JSON(product)
}
func Createproducts(c *fiber.Ctx) error {
    productCollection := db.GetProductCollection()

    // Parse the request body
    var product models.Product
    if err := c.BodyParser(&product); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Could not parse the request body",
        })
    }

    // Validate product data
    if product.Name == "" || product.Price < 0 || product.PID == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid product data",
        })
    }

    // Check if a product with the same PID already exists
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    filter := bson.M{"pid": product.PID}
    var existingProduct models.Product
    err := productCollection.FindOne(ctx, filter).Decode(&existingProduct)

    if err == nil {
        // Product with the same PID already exists
        return c.Status(fiber.StatusConflict).JSON(fiber.Map{
            "error": "Product with the same PID already exists",
        })
    } else if err.Error() != "mongo: no documents in result" {
        // Return error if it's something other than "no documents found"
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Could not check for existing product",
        })
    }

    // Initialize other product details
    product.CreatedAt = time.Now()

    // Insert new product into the database
    _, err = productCollection.InsertOne(ctx, product)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Could not create product",
        })
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Product created successfully",
        "product": product,
    })
}

func GetProductByID(c *fiber.Ctx) error {
	// Get product ID from URL parameters
	productID := c.Params("id")

	objectID, err := primitive.ObjectIDFromHex(productID)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":"Invalid product ID format",
		})
	}

	productCollection := db.GetProductCollection()
	
	filter := bson.M{"_id": objectID}
	var existProduct models.Product
	err = productCollection.FindOne(context.TODO(), filter).Decode(&existProduct)
	if err != nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Product not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(existProduct)
}