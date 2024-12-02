package models

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"mohit.com/ecom-api/db"
)

// User represents the user model for registration and login
type User struct {
	ID        string `json:"id" bson:"_id"`
	Email     string `json:"email" bson:"email"`
	Password  string `json:"password" bson:"password"`
	FirstName string `json:"firstname" bson:"firstname"`
	LastName  string `json:"lastname" bson:"lastname"`
	Role      string `json:"role" bson:"role"` // User role
}

// UserLoginInput represents the input structure for user login
type UserLoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserEmailInput represents the input structure for operations based on email
type UserEmailInput struct {
	Email string `json:"email"`
}

// ResetPasswordInput represents the input structure for password reset
type ResetPasswordInput struct {
	Email       string `json:"email"`
	NewPassword string `json:"new_password"`
}

// FindUserByEmail retrieves a user by their email from the database
func FindUserByEmail(email string) (*User, error) {
	userCollection := db.GetUserCollection()

	var user User
	err := userCollection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

// FindUserByID retrieves a user by their ID from the database
func FindUserByID(id string) (*User, error) {
	userCollection := db.GetUserCollection()

	var user User
	err := userCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}
func (u *User) Save() error {
	userCollection := db.GetUserCollection()

	// Check if the user already exists by email (for updates)
	err := userCollection.FindOne(context.TODO(), bson.M{"email": u.Email}).Err()

	if err == nil {
		// User exists, perform update
		_, err := userCollection.UpdateOne(
			context.TODO(),
			bson.M{"email": u.Email}, // Filter by email
			bson.M{"$set": bson.M{    // Update password, first_name, and last_name
				"password":   u.Password,
				"first_name": u.FirstName,
				"last_name":  u.LastName,
			}},
		)
		return err // Return the update error (if any)
	}

	// If no user is found with the email (mongo.ErrNoDocuments), insert the new user
	if err == mongo.ErrNoDocuments {
		_, err = userCollection.InsertOne(context.TODO(), u)
	}

	return err // Return the insert error (if any)
}
