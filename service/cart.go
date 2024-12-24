package service

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mohit.com/ecom-api/db"
	"mohit.com/ecom-api/models"
)



func AddItemToCart(userID, productID, productName string, quantity int, price float64) error {
	cartCollection := db.GetCartCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Find user's cart
	filter := bson.M{"user_id": userID}
	var cart models.Cart

	err := cartCollection.FindOne(ctx, filter).Decode(&cart)
	if err == mongo.ErrNoDocuments {
		// Create a new cart if it doesn't exist
		cart = models.Cart{
			CartID:     primitive.NewObjectID().Hex(),
			UserID:     userID,
			Items:      []models.CartItem{},
			TotalPrice: 0,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}
	} else if err != nil {
		return err
	}

	// Check if product already exists in the cart
	found := false
	for i, item := range cart.Items {
		if item.ProductID == productID {
			cart.Items[i].Quantity += quantity
			cart.Items[i].Price = price * float64(cart.Items[i].Quantity)
			found = true
			break
		}
	}

	if !found {
		// Add new item to the cart
		cart.Items = append(cart.Items, models.CartItem{
			ProductID:   productID,
			ProductName: productName, // Save the product name here
			Quantity:    quantity,
			Price:       price * float64(quantity),
		})
	}

	// Update the total price
	cart.TotalPrice = 0
	for _, item := range cart.Items {
		cart.TotalPrice += item.Price
	}

	cart.UpdatedAt = time.Now()

	// Upsert (update if exists, insert if not)
	_, err = cartCollection.UpdateOne(ctx, filter, bson.M{"$set": cart}, options.Update().SetUpsert(true))
	return err
}

// GetCart retrieves a user's cart
func GetCart(userID string) (*models.Cart, error) {
	cartCollection := db.GetCartCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var cart models.Cart
	err := cartCollection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&cart)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return &cart, err
}


func RemoveItemFromCart(userID, productID string) error {
	cartCollection := db.GetCartCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Find the user's cart
	filter := bson.M{"user_id": userID}
	var cart models.Cart
	err := cartCollection.FindOne(ctx, filter).Decode(&cart)
	if err == mongo.ErrNoDocuments {
		return fmt.Errorf("cart not found for user: %v", userID)
	} else if err != nil {
		return fmt.Errorf("failed to find cart: %v", err)
	}

	// Find the item to remove
	var itemPrice float64
	for i, item := range cart.Items {
		if item.ProductID == productID {
			// Store the price to update the total price later
			itemPrice = item.Price
			// Remove the item from the cart
			cart.Items = append(cart.Items[:i], cart.Items[i+1:]...)
			break
		}
	}

	// If the item was not found, return an error
	if itemPrice == 0 {
		return fmt.Errorf("item not found in cart: %v", productID)
	}

	// Update the total price
	cart.TotalPrice -= itemPrice

	// Update the cart in the database
	cart.UpdatedAt = time.Now()
	_, err = cartCollection.UpdateOne(ctx, filter, bson.M{"$set": cart})
	if err != nil {
		return fmt.Errorf("failed to update cart: %v", err)
	}
	return nil
}


func ClearCart(userID string) error {
	cartCollection := db.GetCartCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Update the cart by setting the "items" array to an empty array
	_, err := cartCollection.UpdateOne(
		ctx,
		bson.M{"user_id": userID}, // Filter to find the user's cart
		bson.M{"$set": bson.M{
			"items": []interface{}{},
			"total_price": 0,
		}}, // Set "items" to an empty array
	)
	
	return err
}