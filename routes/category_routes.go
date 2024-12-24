package routes

import (
	"github.com/gofiber/fiber/v2"
	categoryController "mohit.com/ecom-api/controllers/category"
	"mohit.com/ecom-api/helper"
)

func CategoryRoutes(app *fiber.App){
		//Category Routes
	categories := app.Group("/api/categories")
	categories.Get("/", helper.Authorize("user"), categoryController.GetAllCategory)
	categories.Post("/", helper.Authorize("admin"),  categoryController.CreateCategory)
	categories.Get("/:id", helper.Authorize("user"), categoryController.GetCategoryByID)
	categories.Put("/:id", helper.Authorize("admin"),  categoryController.UpdateCategory)
	categories.Delete("/:id", helper.Authorize("admin"), categoryController.DeleteCategoryByID)
}