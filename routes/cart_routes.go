package routes

import (
	"github.com/gofiber/fiber/v2"
	"mohit.com/ecom-api/controllers"
	"mohit.com/ecom-api/middleware"
)

func CartRoutes(app *fiber.App) {

	//Cart Routes
	cart := app.Group("/api/cart")
	cart.Get("/", middleware.Authorize("user"), Controllers.GetCart)
	cart.Post("/", middleware.Authorize("user"),  Controllers.AddItemToCart)
	cart.Delete("/", middleware.Authorize("user"),  Controllers.ClearCart)
	cart.Delete("/item/:product_id", middleware.Authorize("user"), Controllers.RemoveItemFromCart)
}