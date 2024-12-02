package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mohit.com/ecom-api/db"
)

// Product structure representing a product in the database
type Product struct {
	PID         string    `bson:"pid,omitempty" json:"pid"`             // Product ID
	Name        string    `bson:"name" json:"name"`                     // Product Name
	Description string    `bson:"description" json:"description"`       // Product Description
	Price       float64   `bson:"price" json:"price"`                   // Price of Product
	CategoryID  primitive.ObjectID    `bson:"category_id" json:"category_id"`       // Category ID (Reference)
	BrandID     primitive.ObjectID    `bson:"brand_id" json:"brand_id"`             // Brand ID (Reference)
	Stock       int       `bson:"stock" json:"stock"`                   // Available stock
	Rating      float64   `bson:"rating,omitempty" json:"rating"`       // Product Rating
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`         // Creation Date
	UpdatedAt   time.Time `bson:"updated_at" json:"updated_at"`         // Last Updated Date
}

// GetAllProducts retrieves all products from the database
func GetAllProducts() ([]Product, error) {
	productCollection := db.GetProductCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := productCollection.Find(ctx, bson.M{}, options.Find())
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []Product
	if err := cursor.All(ctx, &products); err != nil {
		return nil, err
	}

	return products, nil
}

func CheckIfProductExists(productName string, brandID primitive.ObjectID) (bool, error) {
    productCollection := db.GetProductCollection()
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Combine filters with $and operator
    filter := bson.M{
        "$and": []bson.M{
            {"name": productName},
            {"brand_id": brandID},
        },
    }

    var existingProduct Product
    err := productCollection.FindOne(ctx, filter).Decode(&existingProduct)
    
    if err == mongo.ErrNoDocuments {
        return false, nil // No product found
    } else if err != nil {
        return false, err // Return the error if something else goes wrong
    }

    return true, nil // Product exists
}

func CreateProduct(name, description string, price, rating float64, category_id,brand_id primitive.ObjectID,
	stock int) error {

    productCollection := db.GetProductCollection()
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

	newProduct := Product{
		Name:        name,
		Description: description,
		Price:       price,
		Rating: rating,
		CategoryID: category_id,
		BrandID: brand_id,
		Stock: stock,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}


    // Attempt to insert the new product
    _, err := productCollection.InsertOne(ctx, newProduct)
    if err != nil {
        // Handle the case when the product already exists
        if err.Error() == "mongo: duplicate key error" {
            return &ProductConflictError{"Product with the same PID already exists"}
        }
        return err // Return other errors if any
    }

    return nil // Successfully created the product
}


// GetProductByPID retrieves a product by its PID
func GetProductByPID(pid string) (*Product, error) {
	productCollection := db.GetProductCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"pid": pid}
	var product Product
	err := productCollection.FindOne(ctx, filter).Decode(&product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

// UpdateProductByPID updates a product by its PID
func UpdateProductByPID(pid string, updatedProduct Product) error {
	productCollection := db.GetProductCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Set the updated time
	updatedProduct.UpdatedAt = time.Now()

	filter := bson.M{"pid": pid}
	update := bson.M{"$set": updatedProduct}

	result, err := productCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return &ProductNotFoundError{"Product not found"}
	}

	return nil
}

// DeleteProductByPID deletes a product by its PID
func DeleteProductByPID(pid string) error {
	productCollection := db.GetProductCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"pid": pid}
	result, err := productCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return &ProductNotFoundError{"Product not found"}
	}

	return nil
}

// Custom Error Types
type ProductConflictError struct{ 
	Message string 
}
func (e *ProductConflictError) Error() string { 
	return e.Message 
}

type ProductNotFoundError struct{ Message string }
func (e *ProductNotFoundError) Error() string { return e.Message }
