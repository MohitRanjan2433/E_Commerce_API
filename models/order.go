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

type Order struct {
	UserID     string       `bson:"user_id" json:"user_id"`             // User ID
	Items      []OrderItem  `bson:"items" json:"items"`                 // Ordered Items
	TotalPrice float64      `bson:"total_price" json:"total_price"`     // Total Price
	Status     string       `bson:"status" json:"status"`               // Order Status (e.g., Pending, Shipped, Delivered)
	Shipping   ShippingInfo `bson:"shipping" json:"shipping"`           // Shipping Information
	PlacedAt   time.Time    `bson:"placed_at" json:"placed_at"`         // Order Placement Date
	UpdatedAt  time.Time    `bson:"updated_at" json:"updated_at"`       // Last Updated Date
}

type OrderItem struct {
	ProductID  primitive.ObjectID  `bson:"product_id" json:"product_id"` // Product ID
	Quantity  int     `bson:"quantity" json:"quantity"`     // Quantity of Product
	Price     float64 `bson:"price" json:"price"`           // Price of Product at the time of order
}

type ShippingInfo struct {
	Address string `bson:"address" json:"address"`
	City    string `bson:"city" json:"city"`
	State   string `bson:"state" json:"state"`
	ZipCode string `bson:"zip_code" json:"zip_code"`
	Country string `bson:"country" json:"country"`
}

type CreateOrderRequest struct {
	UserID       string              `json:"user_id"`       // User ID
	Items        []OrderItem         `json:"items"`         // List of items
	ShippingInfo ShippingInfo        `json:"shipping_info"` // Shipping Information
}

func CreateOrder(userID string, items []OrderItem, totalPrice float64, shippingInfo ShippingInfo) (*Order, error) {
	orderCollection := db.GetOrderCollection()

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	newOrder := Order{
		UserID:     userID,
		Items:      items,
		TotalPrice: totalPrice,
		Status:     "Pending",
		Shipping:   shippingInfo,
		PlacedAt:   time.Now(),
		UpdatedAt:  time.Now(),
	}

	_, err := orderCollection.InsertOne(ctx, newOrder)
	if err != nil {
		return nil, errors.New("Failed to create order: " + err.Error())
	}

	return &newOrder, nil
}


func GetAllOrdersById(userID string) ([]Order, error){
	orderCollection := db.GetOrderCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"user_id": userID}
	cursor, err := orderCollection.Find(ctx, filter, options.Find()); if err != nil{
		return nil, errors.New("Failed to get orders: " + err.Error())
	}
	defer cursor.Close(ctx)

	var orders []Order
	if err := cursor.All(ctx , &orders); err != nil{
		return nil, errors.New("Failed to get orders: " + err.Error())
	}

	return orders, nil
}

func UpdateOrderStatus(newStatus string, orderID primitive.ObjectID) error {
	orderCollection := db.GetOrderCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": orderID}
	update := bson.M{
		"$set": bson.M{
			"status": newStatus,
			"updated_at": time.Now(),
		},
	}

	result, err := orderCollection.UpdateOne(ctx, filter, update)
	if err != nil{
		return errors.New("Failed to update order status: " + err.Error())
	}

	if result.MatchedCount == 0{
		return errors.New("Order not found")
	}

	return nil
}

