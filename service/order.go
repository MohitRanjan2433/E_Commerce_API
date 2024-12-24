package service

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mohit.com/ecom-api/db"
	"mohit.com/ecom-api/models"
)


func CreateOrder(userID string, items []models.OrderItem, totalPrice float64, shippingInfo models.ShippingInfo) (*models.Order, error) {
	orderCollection := db.GetOrderCollection()
	productCollection := db.GetProductCollection()
	inventoryCollection := db.GetInventoryCollection()

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	newOrder := models.Order{
		UserID:     userID,
		Items:      items,
		TotalPrice: totalPrice,
		Status:     "Pending",
		Shipping:   shippingInfo,
		PlacedAt:   time.Now(),
		UpdatedAt:  time.Now(),
	}



	for _, items := range items{
		productID := items.ProductID

		var productRecord models.Product
		err := productCollection.FindOne(ctx, bson.M{"_id": productID}).Decode(&productRecord)
		if err != nil{
			if err == mongo.ErrNoDocuments{
				return nil, errors.New("product not found")
			}
			return nil, errors.New("failed to check product in productDB" + err.Error())
		}

		productFilter := bson.M{"_id": productID}
		productUpdate := bson.M{
			"$inc": bson.M{
				"stock": -items.Quantity,
			},
			"$set": bson.M{
				"updated_at": time.Now(),
			},
		}

		_, err = productCollection.UpdateOne(ctx, productFilter, productUpdate)
		if err != nil{
			return nil, errors.New("Failed to updated product stock" + err.Error())
		}
	}

	for _, items := range items{
		productID := items.ProductID

		productFilter := bson.M{"product_id": productID}
		productUpdate := bson.M{
			"$inc": bson.M{
				"stock": -items.Quantity,
			},
			"$set": bson.M{
				"updated_at": time.Now(),
			},
		}

		_, err := inventoryCollection.UpdateOne(ctx, productFilter, productUpdate)
		if err != nil{
			return nil, errors.New("Failed to updated inventory stock" + err.Error())
		}
	}

	_, err := orderCollection.InsertOne(ctx, newOrder)
	if err != nil {
		return nil, errors.New("Failed to create order: " + err.Error())
	}

	return &newOrder, nil
}


func GetAllOrdersById(userID string) ([]models.Order, error){
	orderCollection := db.GetOrderCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"user_id": userID}
	cursor, err := orderCollection.Find(ctx, filter, options.Find()); if err != nil{
		return nil, errors.New("Failed to get orders: " + err.Error())
	}
	defer cursor.Close(ctx)

	var orders []models.Order
	if err := cursor.All(ctx , &orders); err != nil{
		return nil, errors.New("Failed to get orders: " + err.Error())
	}

	return orders, nil
}

func UpdateOrderStatus(newStatus string, orderID primitive.ObjectID) error {
	orderCollection := db.GetOrderCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": orderID}
	update := bson.M{
		"$set": bson.M{
			"status": newStatus,
			"updated_at": time.Now(),
		},
	}

	result, err := orderCollection.UpdateOne(ctx, filter, update)
	if err != nil{
		return errors.New("Failed to update order status: " + err.Error())
	}

	if result.MatchedCount == 0{
		return errors.New("Order not found")
	}

	return nil
}

