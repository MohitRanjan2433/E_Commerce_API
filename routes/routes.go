package routes

import (
	"github.com/gofiber/fiber/v2"
	"mohit.com/ecom-api/controllers"
	"mohit.com/ecom-api/middleware"
)

func SetupRoutes(app *fiber.App) {
	// Auth Routes
	auth := app.Group("/api/auth")
	auth.Post("/signup", Controllers.Signup)
	auth.Post("/login", Controllers.Login)
	auth.Get("/me", middleware.Authorize("user"),  Controllers.Me)

	// Product Routes
	products := app.Group("/api/products")
	products.Get("/", Controllers.GetAllProducts)
	products.Post("/", middleware.Authorize("admin"), Controllers.CreateProduct)
	products.Get("/:id", Controllers.GetProductByID)
	products.Put("/:id", middleware.Authorize("admin"), Controllers.UpdateProductByID)
	products.Delete("/:id", middleware.Authorize("admin"), Controllers.DeleteProductByID)

	//Category Routes
	categories := app.Group("/api/categories")
	categories.Get("/", Controllers.GetAllCategory)
	categories.Post("/",  Controllers.CreateCategory)
	categories.Get("/:id", Controllers.GetCategoryByID)
	categories.Put("/:id",  Controllers.UpdateCategory)
	categories.Delete("/:id", Controllers.DeleteCategoryByID)

	//Brand Routes
	brands := app.Group("/api/brands")
	brands.Get("/:id", Controllers.GetBrandController)
	brands.Post("/",  Controllers.CreateBrandController)
	brands.Get("/", Controllers.GetAllBrandController)
	brands.Put("/:id", Controllers.UpdateBrandController)
	brands.Delete("/:id", Controllers.DeleteBrandController)

	//Cart Routes
	cart := app.Group("/api/cart")
	cart.Get("/", middleware.Authorize("user"), Controllers.GetCart)
	cart.Post("/", middleware.Authorize("user"),  Controllers.AddItemToCart)
	cart.Delete("/", middleware.Authorize("user"),  Controllers.ClearCart)
	cart.Delete("/item/:product_id", middleware.Authorize("user"), Controllers.RemoveItemFromCart)

	//Order Routes
	orders := app.Group("/api/orders")
	orders.Get("/", middleware.Authorize("user"), Controllers.GetAllOrdersController)
	orders.Post("/", middleware.Authorize("user"), Controllers.CreateOrderHandler)
	// orders.Get("/:id", middleware.IsAuthenticated, controllers.GetOrderByID)
	orders.Put("/status", Controllers.UpdateOrderStatusController)
	// orders.Delete("/:id", middleware.IsAdmin, controllers.DeleteOrder)

	//Inventory Routes
	inventory := app.Group("/api/inventory")
	inventory.Post("/",  Controllers.CreateInventoryController)
	inventory.Get("/",  Controllers.GetAllInventoryController)
	inventory.Post("/updateStock",  Controllers.UpdateStockController)
}

