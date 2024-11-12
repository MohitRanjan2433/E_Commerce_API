package middleware

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"mohit.com/ecom-api/db"
	"mohit.com/ecom-api/models"
)
func IsAuthenticated(c *fiber.Ctx) error {
	// Step 1: Get the Authorization header
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authorization header is required. Please provide a valid token in the Authorization header.",
		})
	}

	// Step 2: Split the token from the "Bearer " prefix
	tokenParts := strings.Split(authHeader, "Bearer ")
	if len(tokenParts) < 2 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid Authorization header format. The correct format is 'Bearer <token>'.",
		})
	}
	tokenString := tokenParts[1]

	// Step 3: Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	// Step 4: Token validation
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": fmt.Sprintf("Invalid token: %v", err.Error()),
		})
	}
	if !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "The token is expired or invalid.",
		})
	}

	// Step 5: Extract claims and validate
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Failed to parse token claims. The token may be malformed.",
		})
	}

	email, ok := claims["email"].(string)
	if !ok || email == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Email is missing from token claims or is empty. The token may be corrupted.",
		})
	}

	// Step 6: Check if user exists in the database
	userCollection := db.GetUserCollection()
	var user models.User
	filter := bson.M{"email": email}
	err = userCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "User not found. Please check the email in the token or register the user.",
		})
	}

	// Store user information in the context for further use
	c.Locals("user", user)
	return nil // Return nil when authentication succeeds
}

func IsAdmin(c *fiber.Ctx) error {
	// Ensure that the user is authenticated
	if err := IsAuthenticated(c); err != nil {
		return err // Return the error directly if authentication fails
	}

	// Retrieve the user from the context
	user, ok := c.Locals("user").(models.User)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve user information from the context.",
		})
	}

	// Check if the user has an admin role
	if user.Role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "You do not have permission to access this resource. Admin privileges are required.",
		})
	}

	// Allow the request to proceed to the next handler
	return c.Next()
}
