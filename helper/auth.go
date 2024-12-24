package helper

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


func Authorize(requiredRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {

		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authorization header is required.",
			})
		}

		tokenParts := strings.Split(authHeader, "Bearer ")
		if len(tokenParts) < 2 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid Authorization header format. Use 'Bearer <token>'.",
			})
		}
		tokenString := tokenParts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or expired token.",
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Failed to parse token claims.",
			})
		}

		email, ok := claims["email"].(string)
		if !ok || email == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Email missing in token claims.",
			})
		}

		userCollection := db.GetUserCollection()
		var user models.User
		filter := bson.M{"email": email}
		err = userCollection.FindOne(context.TODO(), filter).Decode(&user)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "User not found.",
			})
		}

		if len(requiredRoles) > 0 {
			isAuthorized := false
			for _, role := range requiredRoles {
				if user.Role == role {
					isAuthorized = true
					break
				}
			}
			if !isAuthorized {
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
					"error": "You do not have the required permissions.",
				})
			}
		}

		c.Locals("user", user)
		c.Locals("userID", user.ID)

		return c.Next()
	}
}

