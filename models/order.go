package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
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
