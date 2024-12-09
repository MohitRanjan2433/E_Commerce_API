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

