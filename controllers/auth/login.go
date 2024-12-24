package auth

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"mohit.com/ecom-api/db"
	"mohit.com/ecom-api/middleware"
	"mohit.com/ecom-api/models"
)

func Login(c *fiber.Ctx) error{
	var loginInput models.User
	if err := c.BodyParser(&loginInput); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request data",
		})
	}

	userCollection := db.GetUserCollection()

	var user models.User
	filter := bson.M{"email": loginInput.Email}
	err := userCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil{
		if err == mongo.ErrNoDocuments{
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid email or password",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error finding user in database",
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInput.Password))
	if err != nil{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	token, err := middleware.GenerateJWT(user)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error generating token",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

// func GenerateJWT(user models.User) (string, error){
// 	claims := jwt.MapClaims{
// 		"email": user.Email,
// 		"role": user.Role,
// 		"exp" : time.Now().Add(time.Hour * 72).Unix(),
// 	}

// 	token :=jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
// }

func Me(c *fiber.Ctx) error{
	user := c.Locals("user")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user": user,
	})
}