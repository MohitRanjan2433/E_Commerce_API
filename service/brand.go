package service

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mohit.com/ecom-api/db"
	"mohit.com/ecom-api/models"
)

func CheckIfBrandExists(brandName string) (bool, error){
	brandCollection := db.GetBrandCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"name": brandName}
	var existingBrand models.Brand
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

	newBrand := models.Brand{
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

func GetBrandByID(brandID primitive.ObjectID) (*models.Brand , error){

	brandCollection := db.GetBrandCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"brand_id": brandID}
	var brand models.Brand
	err := brandCollection.FindOne(ctx, filter).Decode(&brand)
	if err == mongo.ErrNoDocuments{
		return nil, nil
	}else if err != nil{
		return nil, err
	}

	return &brand, nil

}

func DeleteBrand(brandID primitive.ObjectID) error {
	brandCollection := db.GetBrandCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"brand_id": brandID}
	_, err := brandCollection.DeleteOne(ctx, filter)
	if err != nil{
		return err
	}

	return nil

}

func UpdateBrand(brandID primitive.ObjectID, name ,country string) error {
	brandCollection := db.GetBrandCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"brand_id": brandID}
	update := bson.M{
		"$set": bson.M{
			"name": name,
			"country": country,
			"updated_at": time.Now(),
		},
	}

	_, err := brandCollection.UpdateOne(ctx, filter, update)
	if err != nil{
		return err
	}
	return nil
}

func GetAllBrand() ([]models.Brand, error){
	brandCollection := db.GetBrandCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := brandCollection.Find(ctx, bson.M{}, options.Find()); if err != nil{
		return nil, err
	}

	defer cursor.Close(ctx)
	var brand []models.Brand

	if err := cursor.All(ctx, &brand); err != nil{
		return nil, err
	}

	return brand, nil
}