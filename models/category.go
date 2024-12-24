package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

)

type Category struct {
	CategoryID primitive.ObjectID `bson:"category_id,omitempty" json:"category_id"` // Category ID
	Name       string        `bson:"name" json:"name"`                         // Category Name
	CreatedAt  time.Time     `bson:"created_at" json:"created_at"`             // Creation Date
	UpdatedAt  time.Time     `bson:"updated_at" json:"updated_at"`             // Last Updated Date
}


