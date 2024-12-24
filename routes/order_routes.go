package routes

import (
	"github.com/gofiber/fiber/v2"
	orderController "mohit.com/ecom-api/controllers/order"
	"mohit.com/ecom-api/helper"
)

func OrderRoutes(app *fiber.App) {

	//Order Routes
	orders := app.Group("/api/orders")
	orders.Get("/", helper.Authorize("user"), orderController.GetAllOrdersController)
	orders.Post("/", helper.Authorize("user"), orderController.CreateOrderHandler)
	// orders.Get("/:id", middleware.IsAuthenticated, controllers.GetOrderByID)
	orders.Put("/status", orderController.UpdateOrderStatusController)
	// orders.Delete("/:id", middleware.IsAdmin, controllers.DeleteOrder)
}