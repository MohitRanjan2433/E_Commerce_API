package models

import "time"

type Inventory struct {
	ProductID string    `bson:"product_id,omitempty" json:"product_id"` // Product ID
	Stock     int       `bson:"stock" json:"stock"`                     // Quantity in Stock
	Warehouse string    `bson:"warehouse" json:"warehouse"`             // Warehouse Location
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`           // Last Updated Date
}
