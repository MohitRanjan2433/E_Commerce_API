package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mohit.com/ecom-api/db"
)

type Category struct {
	CategoryID primitive.ObjectID `bson:"category_id,omitempty" json:"category_id"` // Category ID
	Name       string        `bson:"name" json:"name"`                         // Category Name
	CreatedAt  time.Time     `bson:"created_at" json:"created_at"`             // Creation Date
	UpdatedAt  time.Time     `bson:"updated_at" json:"updated_at"`             // Last Updated Date
}


func CheckIfCategoryExists(categoryName string) (bool, error) {
    categoryCollection := db.GetCategoryCollection()
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    filter := bson.M{"name": categoryName}
    var existingCategory Category

    // Try to find a document with the given category name
    err := categoryCollection.FindOne(ctx, filter).Decode(&existingCategory)

    if err == mongo.ErrNoDocuments {
        // If no category is found, return false (no conflict), and no error
        return false, nil
    } else if err != nil {
        // If there's any other error, return the error
        return false, err
    }

    // If a category is found, return true (category exists)
    return true, nil
}



func CreateCategory(categoryName string) (primitive.ObjectID, error){
	categoryCollection := db.GetCategoryCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	newCategory := Category{
		CategoryID: primitive.NewObjectID(),
		Name: categoryName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := categoryCollection.InsertOne(ctx, newCategory)
	if err != nil{
		return primitive.NilObjectID, err
	}

	return primitive.NewObjectID(), nil
}

func GetAllCategory() ([]Category, error){
	categoryCollection := db.GetCategoryCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := categoryCollection.Find(ctx, bson.M{}, options.Find()); if err != nil{
		return nil, err
	}

	defer cursor.Close(ctx)
	var category []Category
	if err := cursor.All(ctx, &category); err != nil{
		return nil, err
	}

	return category, nil
}