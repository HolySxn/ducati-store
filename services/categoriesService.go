package services

import (
	"context"
	"ducati-store/database"
	"ducati-store/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CategoryInput represents the input data for creating or updating a category
type CategoryInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// GetAllCategories retrieves all categories from the database
func GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	collection := database.DB.Collection("categories")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var category models.Category
		cursor.Decode(&category)
		categories = append(categories, category)
	}

	return categories, nil
}

// CreateCategory adds a new category to the database
func CreateCategory(input CategoryInput) (primitive.ObjectID, error) {
	id := primitive.NewObjectID()
	category := models.Category{
		ID:          id,
		Name:        input.Name,
		Description: input.Description,
	}

	collection := database.DB.Collection("categories")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, category)
	return id, err
}

// RemoveCategory deletes a category from the database
func RemoveCategory(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	collection := database.DB.Collection("categories")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}

// UpdateCategory modifies a category in the database
func UpdateCategory(id string, input CategoryInput) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	collection := database.DB.Collection("categories")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"name":        input.Name,
			"description": input.Description,
		},
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	return err
}

// GetCategoryByID retrieves a category by its ID
func GetCategoryByID(id string) (models.Category, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Category{}, err
	}

	var category models.Category
	collection := database.DB.Collection("categories")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&category)
	return category, err
}
