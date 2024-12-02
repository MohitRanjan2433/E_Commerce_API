package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"mohit.com/ecom-api/db"
)

type Brand struct {
	BrandID   primitive.ObjectID `bson:"brand_id,omitempty" json:"brand_id"` // Brand ID
	Name      string       `bson:"name" json:"name"`                   // Brand Name
	Country   string       `bson:"country" json:"country"`             // Country of Origin
	CreatedAt time.Time    `bson:"created_at" json:"created_at"`       // Creation Date
	UpdatedAt time.Time    `bson:"updated_at" json:"updated_at"`       // Last Updated Date
}


func CheckIfBrandExists(brandName string) (bool, error){
	brandCollection := db.GetBrandCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"name": brandName}
	var existingBrand Brand
	err := brandCollection.FindOne(ctx, filter).Decode(&existingBrand)

	if err == mongo.ErrNoDocuments{
		return false, nil
	}else if err!= nil{
		return false, err
	}

	return true, nil
}

func CreateBrand(brandName ,Country string) (primitive.ObjectID, error){
	brandCollection := db.GetBrandCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	newBrand := Brand{
		BrandID: primitive.NewObjectID(),
		Name:      brandName,
		Country:   Country,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := brandCollection.InsertOne(ctx, newBrand)
	if err != nil{
		return primitive.NilObjectID, err
	}

	return primitive.NewObjectID(), nil
}