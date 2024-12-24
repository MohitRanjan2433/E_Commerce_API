package routes

import (
	"github.com/gofiber/fiber/v2"
	cartController "mohit.com/ecom-api/controllers/cart"
	"mohit.com/ecom-api/helper"
)

func CartRoutes(app *fiber.App) {

	//Cart Routes
	cart := app.Group("/api/cart")
	cart.Get("/", helper.Authorize("user"), cartController.GetCart)
	cart.Post("/", helper.Authorize("user"),  cartController.AddItemToCart)
	cart.Delete("/", helper.Authorize("user"),  cartController.ClearCart)
	cart.Delete("/item/:product_id", helper.Authorize("user"), cartController.RemoveItemFromCart)
}