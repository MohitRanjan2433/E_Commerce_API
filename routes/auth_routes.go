package routes

import (
	"github.com/gofiber/fiber/v2"
	"mohit.com/ecom-api/controllers"
	"mohit.com/ecom-api/middleware"
)

func AuthRoutes(app *fiber.App) {
	// Auth Routes
	auth := app.Group("/api/auth")
	auth.Post("/signup", Controllers.Signup)
	auth.Post("/login", Controllers.Login)
	auth.Get("/me", middleware.Authorize("user"),  Controllers.Me)
}