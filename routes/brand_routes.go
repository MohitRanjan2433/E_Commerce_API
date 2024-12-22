package routes

import (
	"github.com/gofiber/fiber/v2"
	"mohit.com/ecom-api/controllers"
	"mohit.com/ecom-api/middleware"
)

func BrandRoutes(app *fiber.App) {

	//Brand Routes
	brands := app.Group("/api/brands")
	brands.Get("/:id", middleware.Authorize("user"), Controllers.GetBrandController)
	brands.Post("/", middleware.Authorize("admin"), Controllers.CreateBrandController)
	brands.Get("/",middleware.Authorize("user"), Controllers.GetAllBrandController)
	brands.Put("/:id", middleware.Authorize("admin"), Controllers.UpdateBrandController)
	brands.Delete("/:id", middleware.Authorize("admin"), Controllers.DeleteBrandController)
}