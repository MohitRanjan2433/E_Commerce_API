package models

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mohit.com/ecom-api/db"
)

type Inventory struct {
	ProductID string    `bson:"product_id,omitempty" json:"product_id"` // Product ID
	Stock     int       `bson:"stock" json:"stock"`                     // Quantity in Stock
	Warehouse string    `bson:"warehouse" json:"warehouse"`             // Warehouse Location
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`           // Last Updated Date
}


func CreateInventory(userID, productID, warehouse string, stock int) error {
	inventoryCollection := db.GetInventoryCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	inventory := Inventory{
		ProductID: productID,
		Stock: stock,
		Warehouse: warehouse,
		UpdatedAt: time.Now(),
	}

	_, err := inventoryCollection.InsertOne(ctx, inventory)
	if err != nil{
		return errors.New("Error in creating the inventory" + err.Error())
	}

	return nil
}

func GetAllInventory(userID string) ([]Inventory, error){
	inventoryCollection := db.GetInventoryCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := inventoryCollection.Find(ctx, bson.M{}, options.Find()); if err != nil{
		return nil, errors.New("Error in finding the inventoryColl");
	}

	defer cursor.Close(ctx)

	var inventory []Inventory

	if err := cursor.All(ctx, &inventory); err != nil{
		return nil, err
	}

	return inventory, nil
}

func UpdateStock(userID, productID string, pid primitive.ObjectID, stock int) error {
	inventoryCollection := db.GetInventoryCollection()
	productCollection := db.GetProductCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter1 := bson.M{"product_id": productID}
	filter2 := bson.M{"_id": pid}

	update1 := bson.M{
		"$set": bson.M{
			"stock": stock,
			"updated_at": time.Now(),
		},
	}

	update2 := bson.M{
		"$set": bson.M{
			"stock": stock,
		},
	}

	result1, err := inventoryCollection.UpdateOne(ctx, filter1, update1)
	if err != nil{
		return errors.New("Faild to Update stock" + err.Error())
	}

	result2, err := productCollection.UpdateOne(ctx, filter2, update2)
	if err != nil{
		return errors.New("Faild to Update Product stock" + err.Error())
	}

	if result1.MatchedCount == 0 {
		return errors.New("Product not found in inventory coll")
	}

	if result2.MatchedCount == 0 {
		return errors.New("Product not found in produc coll")
	}

	return nil
}