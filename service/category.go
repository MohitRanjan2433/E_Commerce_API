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

func CheckIfCategoryExists(categoryName string) (bool, error) {
    categoryCollection := db.GetCategoryCollection()
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    filter := bson.M{"name": categoryName}
    var existingCategory models.Category

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

	newCategory := models.Category{
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

func GetAllCategory() ([]models.Category, error){
	categoryCollection := db.GetCategoryCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := categoryCollection.Find(ctx, bson.M{}, options.Find()); if err != nil{
		return nil, err
	}

	defer cursor.Close(ctx)
	var category []models.Category
	if err := cursor.All(ctx, &category); err != nil{
		return nil, err
	}

	return category, nil
}

func GetCategoryByID(categoryID string) (*models.Category, error) {
	categoryCollection := db.GetCategoryCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(categoryID)
	filter := bson.M{"category_id": objectID}

	var category models.Category
	err = categoryCollection.FindOne(ctx, filter).Decode(&category)
	if err == mongo.ErrNoDocuments{
		return nil, nil
	}

	return &category, err

}

func UpdateCategory(categoryID , newName string) error {
	categoryCollection := db.GetCategoryCollection()

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	ObjectID, err := primitive.ObjectIDFromHex(categoryID)
	if err != nil {
		return err
	}

	filter := bson.M{"category_id": ObjectID}
	update := bson.M{
		"$set":bson.M{
			"name": newName,
			"updated_at": time.Now(),
		},
	}

	_, err = categoryCollection.UpdateOne(ctx, filter, update)
	return err
}

func DeleteCategoryByID(categoryID string) error {
	categoryCollection := db.GetCategoryCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Convert the categoryID to an ObjectID
	objectID, err := primitive.ObjectIDFromHex(categoryID)
	if err != nil {
		return err 
	}

	// Define the filter to find the category
	filter := bson.M{"category_id": objectID}

	// Perform the delete operation
	result, err := categoryCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err 
	}

	// Check if any document was deleted
	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments 
	}

	return nil
}
