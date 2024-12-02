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
	auth.Get("/me", middleware.IsAuthenticated, Controllers.Me)

	// Product Routes
	products := app.Group("/api/products")
	products.Get("/", Controllers.GetAllProducts)
	products.Post("/", middleware.IsAdmin, Controllers.CreateProduct)
	products.Get("/:pid", Controllers.GetProductByPID)
	// products.Put("/:id", middleware.IsAdmin, Controllers.UpdateProduct)
	products.Delete("/:pid", middleware.IsAdmin, Controllers.DeleteProductByPID)

	// // Category Routes
	categories := app.Group("/api/categories")

	categories.Get("/", Controllers.GetAllCategory)
	categories.Post("/",  Controllers.CreateCategory)
	// categories.Get("/:id", controllers.GetCategoryByID)
	// categories.Put("/:id", middleware.IsAdmin, controllers.UpdateCategory)
	// categories.Delete("/:id", middleware.IsAdmin, controllers.DeleteCategory)

	// // Brand Routes
	brands := app.Group("/api/brands")
	// brands.Get("/", controllers.GetBrands)
	brands.Post("/",  Controllers.CreateBrand)
	// brands.Get("/:id", controllers.GetBrandByID)
	// brands.Put("/:id", middleware.IsAdmin, controllers.UpdateBrand)
	// brands.Delete("/:id", middleware.IsAdmin, controllers.DeleteBrand)

	// // Cart Routes
	// cart := app.Group("/api/cart")
	// cart.Get("/", middleware.IsAuthenticated, controllers.GetCart)
	// cart.Post("/", middleware.IsAuthenticated, controllers.AddToCart)
	// cart.Put("/:itemId", middleware.IsAuthenticated, controllers.UpdateCartItem)
	// cart.Delete("/:itemId", middleware.IsAuthenticated, controllers.RemoveFromCart)

	// // Order Routes
	// orders := app.Group("/api/orders")
	// orders.Get("/", middleware.IsAuthenticated, controllers.GetOrders)
	// orders.Post("/", middleware.IsAuthenticated, controllers.CreateOrder)
	// orders.Get("/:id", middleware.IsAuthenticated, controllers.GetOrderByID)
	// orders.Put("/:id", middleware.IsAdmin, controllers.UpdateOrderStatus)
	// orders.Delete("/:id", middleware.IsAdmin, controllers.DeleteOrder)
	// orders.Get("/:id/payment-status", middleware.IsAuthenticated, controllers.GetPaymentStatus)

	// // Inventory Routes
	// inventory := app.Group("/api/inventory")
	// inventory.Get("/", middleware.IsAdmin, controllers.GetInventory)
	// inventory.Get("/:productId", middleware.IsAdmin, controllers.GetProductInventory)
	// inventory.Post("/", middleware.IsAdmin, controllers.UpdateInventory)
	// inventory.Post("/alerts", middleware.IsAdmin, controllers.SetLowStockAlert)

	// // Search Routes
	// search := app.Group("/api/search")
	// search.Get("/products", controllers.SearchProducts)
	// search.Get("/categories", controllers.SearchCategories)
	// search.Get("/brands", controllers.SearchBrands)

	// // Payment Routes
	// payment := app.Group("/api/payment")
	// payment.Post("/stripe/create-checkout-session", controllers.CreateStripeCheckoutSession)
	// payment.Post("/paypal/create-order", controllers.CreatePaypalOrder)
	// payment.Post("/stripe/webhook", controllers.StripeWebhook)
	// payment.Post("/paypal/webhook", controllers.PaypalWebhook)

	// // Shipping Routes
	// shipping := app.Group("/api/shipping")
	// shipping.Post("/address", middleware.IsAuthenticated, controllers.AddShippingAddress)
	// shipping.Get("/addresses", middleware.IsAuthenticated, controllers.GetShippingAddresses)
	// shipping.Put("/address/:id", middleware.IsAuthenticated, controllers.UpdateShippingAddress)
	// shipping.Delete("/address/:id", middleware.IsAuthenticated, controllers.DeleteShippingAddress)
	// shipping.Get("/methods", controllers.GetShippingMethods)
	// shipping.Get("/status/:orderId", controllers.GetShippingStatus)
	// shipping.Post("/track", controllers.TrackShipment)

	// // Admin Routes
	// admin := app.Group("/api/admin")
	// admin.Get("/dashboard", middleware.IsAdmin, controllers.GetDashboardOverview)
	// admin.Get("/reports/sales", middleware.IsAdmin, controllers.GetSalesReport)
	// admin.Get("/reports/users", middleware.IsAdmin, controllers.GetUserReport)
	// admin.Get("/reports/products", middleware.IsAdmin, controllers.GetProductReport)
	// admin.Get("/users", middleware.IsAdmin, controllers.GetUsers)
	// admin.Put("/users/:id", middleware.IsAdmin, controllers.UpdateUser)
	// admin.Delete("/users/:id", middleware.IsAdmin, controllers.DeleteUser)
	// admin.Get("/orders", middleware.IsAdmin, controllers.GetAllOrders)
	// admin.Get("/orders/:id", middleware.IsAdmin, controllers.GetOrderDetails)
	// admin.Put("/orders/:id/status", middleware.IsAdmin, controllers.UpdateOrderStatus)
	// admin.Post("/roles", middleware.IsAdmin, controllers.CreateRole)
	// admin.Get("/roles", middleware.IsAdmin, controllers.GetRoles)
	// admin.Put("/roles/:id", middleware.IsAdmin, controllers.UpdateRole)
	// admin.Delete("/roles/:id", middleware.IsAdmin, controllers.DeleteRole)
}

