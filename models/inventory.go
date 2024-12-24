package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Inventory represents the inventory for a product
type Inventory struct {
	ProductID primitive.ObjectID `bson:"product_id,omitempty" json:"product_id"` // Product ID (Reference)
	Stock     int                `bson:"stock" json:"stock"`                     // Quantity in Stock
	Warehouse string             `bson:"warehouse" json:"warehouse"`             // Warehouse Location
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`           // Last Updated Date
}

// StockAlertMessage represents an alert for low stock
type StockAlertMessage struct {
	Message string
}

// Error method for StockAlertMessage to satisfy the error interface
func (e *StockAlertMessage) Error() string {
	return e.Message
}

// CreateInventory creates a new inventory record
