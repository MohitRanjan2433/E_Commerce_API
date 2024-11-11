package Controllers

import (
	"context"
	"math/rand"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"mohit.com/ecom-api/db"
	"mohit.com/ecom-api/models"
)

const letterdigits = "0123456789"

func generateId(length int) string{
	rng := rand.New(rand.NewSource(time.Now().UnixNano())) // Initialize the RNG
	result := make([]byte, length)
	for i := range result{
		result[i] = letterdigits[rng.Intn(len(letterdigits))]
	}
	return string(result)
}

//signup
func Signup(c *fiber.Ctx) error {
	var userInput models.User
	// Parse incoming JSON data into the userInput struct
	if err := c.BodyParser(&userInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request data",
		})
	}

	// Get the user collection from the database
	userCollection := db.GetUserCollection()

	// Check if a user with the same email already exists
	filter := bson.M{"email": userInput.Email}
	var existingUser models.User
	err := userCollection.FindOne(context.TODO(), filter).Decode(&existingUser)
	if err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email already exists",
		})
	} else if err != mongo.ErrNoDocuments {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	// Hash the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error hashing password",
		})
	}

	newUserID := generateId(5)

	filterByID := bson.M{"_id": newUserID}
	var userwithID models.User
	err = userCollection.FindOne(context.TODO(), filterByID).Decode(&userwithID)
	if err == nil{
		newUserID = generateId(5)
	}

	// Prepare the user data to be inserted
	newUser := models.User{
		ID: userInput.FirstName + newUserID,
		Email:    userInput.Email,
		Password: string(hashedPassword),
		Role:     "user",
		FirstName: userInput.FirstName,
		LastName: userInput.LastName,
	}

	// Insert the new user into the database
	_, err = userCollection.InsertOne(context.TODO(), newUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error saving user to database",
		})
	}

	// Return success response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User created successfully",
	})
}

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

	token, err := GenerateJWT(user)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error generating token",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

func GenerateJWT(user models.User) (string, error){
	claims := jwt.MapClaims{
		"email": user.Email,
		"role": user.Role,
		"exp" : time.Now().Add(time.Hour * 72).Unix(),
	}

	token :=jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}