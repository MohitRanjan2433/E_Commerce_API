package service

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mohit.com/ecom-api/db"
	"mohit.com/ecom-api/models"
)

func GetAllProducts() ([]models.Product, error) {
	productCollection := db.GetProductCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := productCollection.Find(ctx, bson.M{}, options.Find())
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []models.Product
	if err := cursor.All(ctx, &products); err != nil {
		return nil, err
	}

	return products, nil
}

// CheckIfProductExists checks if a product exists by name and brand
func CheckIfProductExists(productName string, brandID primitive.ObjectID) (bool, error) {
	productCollection := db.GetProductCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{
		"name":     productName,
		"brand_id": brandID,
	}

	var existingProduct models.Product
	err := productCollection.FindOne(ctx, filter).Decode(&existingProduct)
	if err == mongo.ErrNoDocuments {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

// CreateProduct creates a new product
func CreateProduct(name, description string, price, rating float64, categoryID, brandID primitive.ObjectID, stock int) error {
	productCollection := db.GetProductCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	newProduct := models.Product{
		Name:        name,
		Description: description,
		Price:       price,
		Rating:      rating,
		CategoryID:  categoryID,
		BrandID:     brandID,
		Stock:       0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	_, err := productCollection.InsertOne(ctx, newProduct)
	if err != nil {
		return err
	}

	return nil
}

// GetProductByID retrieves a product by its ID
func GetProductByID(id primitive.ObjectID) (*models.Product, error) {
	productCollection := db.GetProductCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	var product models.Product
	err := productCollection.FindOne(ctx, filter).Decode(&product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

// UpdateProductByID updates a product by its ID
func UpdateProductByID(id primitive.ObjectID, updatedProduct models.Product) error {
	productCollection := db.GetProductCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	updatedProduct.UpdatedAt = time.Now()

	filter := bson.M{"_id": id}
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

// DeleteProductByID deletes a product by its ID
func DeleteProductByID(id primitive.ObjectID) error {
	productCollection := db.GetProductCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
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
type ProductNotFoundError struct {
	Message string
}

func (e *ProductNotFoundError) Error() string {
	return e.Message
}
