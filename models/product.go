package models

import (
	"time"

)

type Product struct {
	PID        	string			   `bson:"pid,omitempty json:"pid"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	Price       float64            `bson:"price" json:"price"`
	Category    string             `bson:"category" json:"category"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
}