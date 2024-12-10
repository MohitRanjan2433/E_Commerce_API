package Controllers

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mohit.com/ecom-api/models"
)

func CreateBrandController (c *fiber.Ctx) error {
	var request struct{
		Name string `json:"name"`
		Country string `json:"country"`
	}

	if err := c.BodyParser(&request); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":"Invalid Body Request",
		})
	}

	if request.Name == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Brand name is required",
        })
    }

	exists, err := models.CheckIfBrandExists(request.Name)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error checking brand existence",
		})
	}

	if exists {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Brand already exists",
		})
	}

	BrandID, err := models.CreateBrand(request.Name, request.Country)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error creating brand",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Brand created successfully",
		"brand_id": BrandID.Hex(),
	})

}

func GetBrandController(c *fiber.Ctx)error {
	brandIDHex := c.Params("id")
	brandID, err := primitive.ObjectIDFromHex(brandIDHex)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":"Invalid brand id",
		})
	}

	brand, err := models.GetBrandByID(brandID)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":"Error retriving brand",
		})
	}

	if brand == nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":"Brand not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"brand": brand,
	})
}

func DeleteBrandController(c *fiber.Ctx) error {
	brandIDHex := c.Params("id")

	brand, err := primitive.ObjectIDFromHex(brandIDHex)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":"Invalid brand id",
		})
	}

	err = models.DeleteBrand(brand)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error deleting brand",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Brand Deleted Successfully",
	})
}

func UpdateBrandController(c *fiber.Ctx) error {
	brandIDHex := c.Params("id")

	brandID, err := primitive.ObjectIDFromHex(brandIDHex)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Brand ID",
		})
	}

	var request struct{
		Name      string    `json:"name"`
		Country   string    `json:"country"`
	}

	if err := c.BodyParser(&request); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":"Invalid request Body",
		})
	}

	err = models.UpdateBrand(brandID, request.Name, request.Country)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":"Error updating brand",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":"Brand Updated Successfully",
	})
}

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