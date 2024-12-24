package main

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"mohit.com/ecom-api/middleware"
	"mohit.com/ecom-api/routes"
)

func main() {
	app := fiber.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" 
	}

	app.Use(middleware.RateLimiterMiddleware(20, 30*time.Second))

	routes.AuthRoutes(app)
	routes.BrandRoutes(app)
	routes.CategoryRoutes(app)
	routes.CartRoutes(app)
	routes.InventoryRoutes(app)
	routes.OrderRoutes(app)
	routes.ProductRoutes(app)

	// Start server
	err = app.Listen(":" + port)
	if err != nil {
		log.Fatal("Error while starting server:", err)
	}
}
