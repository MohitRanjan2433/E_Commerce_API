package routes

import (
	"github.com/gofiber/fiber/v2"
	"mohit.com/ecom-api/controllers"
	"mohit.com/ecom-api/middleware"
)

func InventoryRoutes(app *fiber.App) {

	//Inventory Routes
	inventory := app.Group("/api/inventory")
	inventory.Post("/", middleware.Authorize("admin"),  Controllers.CreateInventoryController)
	inventory.Get("/", middleware.Authorize("admin"),  Controllers.GetAllInventoryController)
	inventory.Post("/updateStock", middleware.Authorize("admin"),  Controllers.UpdateStockController)
}

