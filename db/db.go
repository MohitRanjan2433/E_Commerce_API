package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client         *mongo.Client
	userCollection *mongo.Collection
	productCollection *mongo.Collection
)

func init() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get MongoDB URL from environment variables
	connectionString := os.Getenv("MONGODB_URL")
	if connectionString == "" {
		log.Fatal("MongoDB URL not set in .env file")
	}

	// Set client options and connect to MongoDB
	clientOptions := options.Client().ApplyURI(connectionString)
	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	// Ping the MongoDB server to verify connection
	if err = client.Ping(context.TODO(), nil); err != nil {
		log.Fatalf("Error pinging MongoDB: %v", err)
	}
	fmt.Println("MongoDB connection successful")

	// Initialize the userCollection
	userCollection = client.Database("E-Commerce").Collection("user")
	productCollection = client.Database("E-Commerce").Collection("products")
}

// GetUserCollection returns the user collection
func GetUserCollection() *mongo.Collection {
	return userCollection
}

func GetProductCollection() *mongo.Collection{
	return productCollection
}

// CloseConnection closes the MongoDB connection
func CloseConnection() {
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Fatalf("Error disconnecting from MongoDB: %v", err)
	}
	fmt.Println("MongoDB connection closed")
}
