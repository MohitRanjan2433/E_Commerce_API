package routes

import (
	"github.com/gofiber/fiber/v2"
	brandConroller "mohit.com/ecom-api/controllers/brand"
	"mohit.com/ecom-api/helper"
)

func BrandRoutes(app *fiber.App) {

	//Brand Routes
	brands := app.Group("/api/brands")
	brands.Get("/:id", helper.Authorize("user"), brandConroller.GetBrandController)
	brands.Post("/", helper.Authorize("admin"), brandConroller.CreateBrandController)
	brands.Get("/",helper.Authorize("user"), brandConroller.GetAllBrandController)
	brands.Put("/:id", helper.Authorize("admin"), brandConroller.UpdateBrandController)
	brands.Delete("/:id", helper.Authorize("admin"), brandConroller.DeleteBrandController)
}