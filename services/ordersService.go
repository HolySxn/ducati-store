package services

import (
	"context"
	"ducati-store/database"
	"ducati-store/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// OrderInput represents the input data for creating or updating an order
type OrderInput struct {
	CustomerID string              `json:"customerId"`
	OrderDate  string              `json:"orderDate"`
	Status     string              `json:"status"`
	Items      []OrderItemInput    `json:"items"`
	Payment    models.PaymentInfo  `json:"payment"`
	Shipping   models.ShippingInfo `json:"shipping"`
}

type OrderItemInput struct {
	MotorcycleID string `json:"motorcycleId"`
	Quantity     int    `json:"quantity"`
}

// GetAllOrders retrieves all orders from the database
func GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	collection := database.DB.Collection("orders")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var order models.Order
		cursor.Decode(&order)
		orders = append(orders, order)
	}

	return orders, nil
}

// CreateOrder adds a new order to the database
func CreateOrder(input OrderInput) (primitive.ObjectID, error) {
	customerID, err := primitive.ObjectIDFromHex(input.CustomerID)
	if err != nil {
		return primitive.NilObjectID, err
	}

	var items []models.OrderItem
	for _, item := range input.Items {
		motorcycleID, err := primitive.ObjectIDFromHex(item.MotorcycleID)
		if err != nil {
			return primitive.NilObjectID, err
		}
		items = append(items, models.OrderItem{
			MotorcycleID: motorcycleID,
			Quantity:     item.Quantity,
		})
	}

	id := primitive.NewObjectID()
	order := models.Order{
		ID:         id,
		CustomerID: customerID,
		OrderDate:  input.OrderDate,
		Status:     input.Status,
		Items:      items,
		Payment:    input.Payment,
		Shipping:   input.Shipping,
	}

	collection := database.DB.Collection("orders")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = collection.InsertOne(ctx, order)
	return id, err
}

// RemoveOrder deletes an order from the database
func RemoveOrder(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	collection := database.DB.Collection("orders")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}

// UpdateOrder modifies an order in the database
func UpdateOrder(id string, input OrderInput) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	customerID, err := primitive.ObjectIDFromHex(input.CustomerID)
	if err != nil {
		return err
	}

	var items []models.OrderItem
	for _, item := range input.Items {
		motorcycleID, err := primitive.ObjectIDFromHex(item.MotorcycleID)
		if err != nil {
			return err
		}
		items = append(items, models.OrderItem{
			MotorcycleID: motorcycleID,
			Quantity:     item.Quantity,
		})
	}

	collection := database.DB.Collection("orders")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"customerId": customerID,
			"orderDate":  input.OrderDate,
			"status":     input.Status,
			"items":      items,
			"payment":    input.Payment,
			"shipping":   input.Shipping,
		},
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	return err
}

// GetOrderByID retrieves an order by its ID
func GetOrderByID(id string) (models.Order, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Order{}, err
	}

	var order models.Order
	collection := database.DB.Collection("orders")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&order)
	return order, err
}
