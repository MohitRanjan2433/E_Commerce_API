package routes

import (
	"github.com/gofiber/fiber/v2"
	authController "mohit.com/ecom-api/controllers/auth"
	"mohit.com/ecom-api/helper"
)

func AuthRoutes(app *fiber.App) {
	// Auth Routes
	auth := app.Group("/api/auth")
	auth.Post("/signup", authController.Signup)
	auth.Post("/login", authController.Login)
	auth.Get("/me", helper.Authorize("user"), authController.Me)
}