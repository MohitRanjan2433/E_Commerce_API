package routes

import (
	"github.com/gofiber/fiber/v2"
	inventoryConroller "mohit.com/ecom-api/controllers/inventory"
	"mohit.com/ecom-api/helper"
)

func InventoryRoutes(app *fiber.App) {

	//Inventory Routes
	inventory := app.Group("/api/inventory")
	inventory.Post("/", helper.Authorize("admin"),  inventoryConroller.CreateInventoryController)
	inventory.Get("/", helper.Authorize("admin"),  inventoryConroller.GetAllInventoryController)
	inventory.Post("/updateStock", helper.Authorize("admin"),  inventoryConroller.UpdateStockController)
}

