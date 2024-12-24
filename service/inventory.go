package service

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mohit.com/ecom-api/db"
	"mohit.com/ecom-api/models"
)


func CreateInventory(productID primitive.ObjectID, warehouse string, inventoryStock int) error {
	inventoryCollection := db.GetInventoryCollection()
	productCollection := db.GetProductCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check if inventory already exists for this product
	var existingInventory models.Inventory
	err := inventoryCollection.FindOne(ctx, bson.M{"product_id": productID}).Decode(&existingInventory)
	if err != nil && err != mongo.ErrNoDocuments {
		return errors.New("failed to check for existing inventory: " + err.Error())
	}

	// If inventory already exists, return an error
	if err == nil {
		return errors.New("inventory for this product already exists")
	}

	// Create a new inventory record
	inventory := models.Inventory{
		ProductID: productID,
		Stock:     inventoryStock,
		Warehouse: warehouse,
		UpdatedAt: time.Now(),
	}

	// Insert new inventory record
	_, err = inventoryCollection.InsertOne(ctx, inventory)
	if err != nil {
		return errors.New("failed to create inventory: " + err.Error())
	}

	// Check if the product exists in the product collection
	productFilter := bson.M{"_id": productID}
	productCount, err := productCollection.CountDocuments(ctx, productFilter)
	if err != nil {
		return errors.New("failed to check product existence: " + err.Error())
	}
	if productCount == 0 {
		return errors.New("product not found in product collection")
	}

	// Update product stock by the inventory stock
	productUpdate := bson.M{
		"$inc": bson.M{
			"stock": inventoryStock,
		},
	}
	_, err = productCollection.UpdateOne(ctx, productFilter, productUpdate)
	if err != nil {
		return errors.New("failed to update product stock: " + err.Error())
	}

	return nil
}

// GetAllInventory retrieves all inventory records
func GetAllInventory() ([]models.Inventory, error) {
	inventoryCollection := db.GetInventoryCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := inventoryCollection.Find(ctx, bson.M{}, options.Find())
	if err != nil {
		return nil, errors.New("failed to fetch inventory: " + err.Error())
	}
	defer cursor.Close(ctx)

	var inventories []models.Inventory
	if err := cursor.All(ctx, &inventories); err != nil {
		return nil, errors.New("failed to parse inventory records: " + err.Error())
	}

	return inventories, nil
}

// UpdateStock updates the stock for a specific product in both inventory and product collections
func UpdateStock(productID primitive.ObjectID, inventoryStock int) error {
	inventoryCollection := db.GetInventoryCollection()
	productCollection := db.GetProductCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Update inventory collection
	inventoryFilter := bson.M{"product_id": productID}
	inventoryUpdate := bson.M{
		"$inc": bson.M{
			"stock": inventoryStock, // Increment stock by the specified amount
		},
		"$set": bson.M{
			"updated_at": time.Now(), // Update the updated_at field to the current time
		},
	}

	inventoryResult, err := inventoryCollection.UpdateOne(ctx, inventoryFilter, inventoryUpdate)
	if err != nil {
		return errors.New("failed to update inventory stock: " + err.Error())
	}

	// Ensure that the inventory record was updated
	if inventoryResult.ModifiedCount == 0 {
		return errors.New("no inventory record updated")
	}

	// Update product collection
	productFilter := bson.M{"_id": productID}
	productUpdate := bson.M{
		"$inc": bson.M{
			"stock": inventoryStock, // Increment stock in the product collection
		},
	}

	productResult, err := productCollection.UpdateOne(ctx, productFilter, productUpdate)
	if err != nil {
		return errors.New("failed to update product stock: " + err.Error())
	}

	// Ensure that the product record was updated
	if productResult.ModifiedCount == 0 {
		return errors.New("no product stock updated")
	}

	return nil
}