package helper

import "go.mongodb.org/mongo-driver/bson/primitive"

func ParseObjectID(idStr string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(idStr)
}