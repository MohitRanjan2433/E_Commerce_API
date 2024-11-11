package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"mohit.com/ecom-api/routes" // Import your routes package
)

func main() {
	// Initialize Fiber app
	app := fiber.New()

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get PORT from environment variables or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // default port
	}

	// Set up routes (make sure this function is correctly defined)
	routes.SetupRoutes(app)

	// Start server
	err = app.Listen(":" + port)
	if err != nil {
		log.Fatal("Error while starting server:", err)
	}
}
