package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Brand struct {
	BrandID   primitive.ObjectID `bson:"brand_id,omitempty" json:"brand_id"` 
	Name      string       `bson:"name" json:"name"`                   
	Country   string       `bson:"country" json:"country"`             
	CreatedAt time.Time    `bson:"created_at" json:"created_at"`       
	UpdatedAt time.Time    `bson:"updated_at" json:"updated_at"`       
}


