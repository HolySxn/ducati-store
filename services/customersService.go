package services

import (
	"context"
	"ducati-store/database"
	"ducati-store/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CustomerInput represents the input data for creating or updating a customer
type CustomerInput struct {
	Name         string   `json:"name"`
	Email        string   `json:"email"`
	Addresses    []string `json:"addresses"`
	PhoneNumbers []string `json:"phoneNumbers"`
}

// GetAllCustomers retrieves all customers from the database
func GetAllCustomers() ([]models.Customer, error) {
	var customers []models.Customer
	collection := database.DB.Collection("customers")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var customer models.Customer
		cursor.Decode(&customer)
		customers = append(customers, customer)
	}

	return customers, nil
}

// CreateCustomer adds a new customer to the database
func CreateCustomer(input CustomerInput) (primitive.ObjectID, error) {
	id := primitive.NewObjectID()
	customer := models.Customer{
		ID:           id,
		Name:         input.Name,
		Email:        input.Email,
		Addresses:    input.Addresses,
		PhoneNumbers: input.PhoneNumbers,
	}

	collection := database.DB.Collection("customers")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, customer)
	return id, err
}

// RemoveCustomer deletes a customer from the database
func RemoveCustomer(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	collection := database.DB.Collection("customers")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}

// UpdateCustomer modifies a customer in the database
func UpdateCustomer(id string, input CustomerInput) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	collection := database.DB.Collection("customers")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"name":         input.Name,
			"email":        input.Email,
			"addresses":    input.Addresses,
			"phoneNumbers": input.PhoneNumbers,
		},
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	return err
}

// GetCustomerByID retrieves a customer by its ID
func GetCustomerByID(id string) (models.Customer, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Customer{}, err
	}

	var customer models.Customer
	collection := database.DB.Collection("customers")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&customer)
	return customer, err
}
