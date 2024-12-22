package routes

import (
	"github.com/gofiber/fiber/v2"
	"mohit.com/ecom-api/controllers"
	"mohit.com/ecom-api/middleware"
)

func CategoryRoutes(app *fiber.App){
		//Category Routes
	categories := app.Group("/api/categories")
	categories.Get("/", middleware.Authorize("user"), Controllers.GetAllCategory)
	categories.Post("/", middleware.Authorize("admin"),  Controllers.CreateCategory)
	categories.Get("/:id", middleware.Authorize("user"), Controllers.GetCategoryByID)
	categories.Put("/:id", middleware.Authorize("admin"),  Controllers.UpdateCategory)
	categories.Delete("/:id", middleware.Authorize("admin"), Controllers.DeleteCategoryByID)
}