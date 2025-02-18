package services

import (
	"context"
	"ducati-store/database"
	"ducati-store/models"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MotorcycleInput represents the input data for creating a motorcycle
type MotorcycleInput struct {
	Model          string   `json:"model"`
	Brand          string   `json:"brand"`
	Year           int      `json:"year"`
	Price          float64  `json:"price"`
	Specifications []string `json:"specifications"`
	Status         string   `json:"status"`
	CategoryID     string   `json:"categoryId"`
}

// GetAllMotorcycles retrieves all motorcycles from the database
func GetAllMotorcycles() ([]models.Motorcycle, error) {
	var motorcycles []models.Motorcycle
	collection := database.DB.Collection("motorcycles")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var motorcycle models.Motorcycle
		cursor.Decode(&motorcycle)
		motorcycles = append(motorcycles, motorcycle)
	}

	return motorcycles, nil
}

// CreateMotorcycle adds a new motorcycle to the database
func CreateMotorcycle(input MotorcycleInput) (primitive.ObjectID, error) {
	categoryID, err := primitive.ObjectIDFromHex(input.CategoryID)
	if err != nil {
		return primitive.NilObjectID, err
	}

	id := primitive.NewObjectID()
	motorcycle := models.Motorcycle{
		ID:             id,
		Model:          input.Model,
		Brand:          input.Brand,
		Year:           input.Year,
		Price:          input.Price,
		Specifications: input.Specifications,
		Status:         input.Status,
		CategoryID:     categoryID,
	}

	collection := database.DB.Collection("motorcycles")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = collection.InsertOne(ctx, motorcycle)
	return id, err
}

// RemoveMotorcycle deletes a motorcycle from the database
func RemoveMotorcycle(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	collection := database.DB.Collection("motorcycles")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}

// UpdateMotorcycle modifies a motorcycle in the database
func UpdateMotorcycle(id string, input MotorcycleInput) error {
	motorcycle, _ := GetMotorcycleByID(id)
	if motorcycle == nil {
		return errors.New("motorcycle not found")
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	var categoryID primitive.ObjectID
	if input.CategoryID != "" {
		categoryID, err = primitive.ObjectIDFromHex(input.CategoryID)
		if err != nil {
			return err
		}
	}else{
		categoryID = motorcycle.(models.Motorcycle).CategoryID
	}
	if input.Model == "" {
		input.Model = motorcycle.(models.Motorcycle).Model
	}
	if input.Brand == "" {
		input.Brand = motorcycle.(models.Motorcycle).Brand
	}
	if input.Year == 0 {
		input.Year = motorcycle.(models.Motorcycle).Year
	}
	if input.Price == 0 {
		input.Price = motorcycle.(models.Motorcycle).Price
	}
	if input.Specifications == nil {
		input.Specifications = motorcycle.(models.Motorcycle).Specifications
	}
	if input.Status == "" {
		input.Status = motorcycle.(models.Motorcycle).Status
	}

	collection := database.DB.Collection("motorcycles")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"model":          input.Model,
			"brand":          input.Brand,
			"year":           input.Year,
			"price":          input.Price,
			"specifications": input.Specifications,
			"status":         input.Status,
			"category_id":     categoryID,
		},
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objectID}, update)

	return err
}

// GetMotorcycleByID retrieves a motorcycle by its ID
func GetMotorcycleByID(id string) (interface{}, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Motorcycle{}, err
	}

	var motorcycle models.Motorcycle
	collection := database.DB.Collection("motorcycles")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&motorcycle)
	return motorcycle, err
}
