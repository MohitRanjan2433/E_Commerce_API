package models

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mohit.com/ecom-api/db"
)

type Cart struct {
	CartID     string     `bson:"cart_id,omitempty" json:"cart_id"` // Cart ID
	UserID     string     `bson:"user_id" json:"user_id"`           // User ID
	Items      []CartItem `bson:"items" json:"items"`               // Items in the Cart
	TotalPrice float64    `bson:"total_price" json:"total_price"`   // Total Price of Cart Items
	CreatedAt  time.Time  `bson:"created_at" json:"created_at"`     // Creation Date
	UpdatedAt  time.Time  `bson:"updated_at" json:"updated_at"`     // Last Updated Date
}

type CartItem struct {
	ProductID   string  `bson:"product_id" json:"product_id"`   // Product ID
	ProductName string  `bson:"product_name" json:"product_name"` // Product name
	Quantity    int     `bson:"quantity" json:"quantity"`       // Quantity of Product
	Price       float64 `bson:"price" json:"price"`             // Price of Product at the time of addition
}

// AddItemToCart adds an item to the user's cart
func AddItemToCart(userID, productID, productName string, quantity int, price float64) error {
	cartCollection := db.GetCartCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Find user's cart
	filter := bson.M{"user_id": userID}
	var cart Cart

	err := cartCollection.FindOne(ctx, filter).Decode(&cart)
	if err == mongo.ErrNoDocuments {
		// Create a new cart if it doesn't exist
		cart = Cart{
			CartID:     primitive.NewObjectID().Hex(),
			UserID:     userID,
			Items:      []CartItem{},
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
		cart.Items = append(cart.Items, CartItem{
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
func GetCart(userID string) (*Cart, error) {
	cartCollection := db.GetCartCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var cart Cart
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

	// Check if cart exists for the user
	filter := bson.M{"user_id": userID}
	update := bson.M{
		"$pull": bson.M{
			"items": bson.M{"product_id": productID},
		},
	}

	// Perform the update (remove the item)
	_, err := cartCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to remove item from cart: %v", err)
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
		bson.M{"$set": bson.M{"items": []interface{}{}}}, // Set "items" to an empty array
	)

	return err
}