package routes

import (
	"github.com/gofiber/fiber/v2"
	"mohit.com/ecom-api/controllers"
	"mohit.com/ecom-api/middleware"
)

func OrderRoutes(app *fiber.App) {

	//Order Routes
	orders := app.Group("/api/orders")
	orders.Get("/", middleware.Authorize("user"), Controllers.GetAllOrdersController)
	orders.Post("/", middleware.Authorize("user"), Controllers.CreateOrderHandler)
	// orders.Get("/:id", middleware.IsAuthenticated, controllers.GetOrderByID)
	orders.Put("/status", Controllers.UpdateOrderStatusController)
	// orders.Delete("/:id", middleware.IsAdmin, controllers.DeleteOrder)
}