package routes

import (
	"github.com/gofiber/fiber/v2"
	productController "mohit.com/ecom-api/controllers/product"
	"mohit.com/ecom-api/helper"
)

func ProductRoutes(app *fiber.App) {

	// Product Routes
	products := app.Group("/api/products")
	products.Get("/", helper.Authorize("user"), productController.GetAllProducts)
	products.Post("/", helper.Authorize("admin"), productController.CreateProduct)
	products.Get("/:id", helper.Authorize("user"), productController.GetProductByID)
	products.Put("/:id", helper.Authorize("admin"), productController.UpdateProductByID)
	products.Delete("/:id", helper.Authorize("admin"), productController.DeleteProductByID)
}