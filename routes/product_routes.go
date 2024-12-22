package routes

import (
	"github.com/gofiber/fiber/v2"
	"mohit.com/ecom-api/controllers"
	"mohit.com/ecom-api/middleware"
)

func ProductRoutes(app *fiber.App) {

	// Product Routes
	products := app.Group("/api/products")
	products.Get("/", middleware.Authorize("user"), Controllers.GetAllProducts)
	products.Post("/", middleware.Authorize("admin"), Controllers.CreateProduct)
	products.Get("/:id", middleware.Authorize("user"), Controllers.GetProductByID)
	products.Put("/:id", middleware.Authorize("admin"), Controllers.UpdateProductByID)
	products.Delete("/:id", middleware.Authorize("admin"), Controllers.DeleteProductByID)
}