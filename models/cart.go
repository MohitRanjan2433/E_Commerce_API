package models

import (
	"time"
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

