package auth

import (
	"context"
	"math/rand"
	"time"

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

func Signup(c *fiber.Ctx) error {
	var userInput models.User
	
	if err := c.BodyParser(&userInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request data",
		})
	}

	userCollection := db.GetUserCollection()

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