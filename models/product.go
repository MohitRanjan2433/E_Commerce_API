package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Product structure representing a product in the database
type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`        // MongoDB ID
	Name        string             `bson:"name" json:"name"`               // Product Name
	Description string             `bson:"description" json:"description"` // Product Description
	Price       float64            `bson:"price" json:"price"`             // Price of Product
	CategoryID  primitive.ObjectID `bson:"category_id" json:"category_id"` // Category ID (Reference)
	BrandID     primitive.ObjectID `bson:"brand_id" json:"brand_id"`       // Brand ID (Reference)
	Stock       int                `bson:"stock" json:"stock"`             // Available stock
	Rating      float64            `bson:"rating,omitempty" json:"rating"` // Product Rating
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`   // Creation Date
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`   // Last Updated Date
}

