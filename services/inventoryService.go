package services

import (
    "context"
    "ducati-store/database"
    "ducati-store/models"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

// InventoryInput represents the input data for creating or updating an inventory record
type InventoryInput struct {
    MotorcycleID string `json:"motorcycleId"`
    Quantity     int    `json:"quantity"`
    Location     string `json:"location"`
}

// GetAllInventory retrieves all inventory records from the database
func GetAllInventory() ([]models.Inventory, error) {
    var inventory []models.Inventory
    collection := database.DB.Collection("inventory")

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    for cursor.Next(ctx) {
        var item models.Inventory
        cursor.Decode(&item)
        inventory = append(inventory, item)
    }

    return inventory, nil
}

// CreateInventory adds a new inventory record to the database
func CreateInventory(input InventoryInput) (primitive.ObjectID, error) {
    motorcycleID, err := primitive.ObjectIDFromHex(input.MotorcycleID)
    if err != nil {
        return primitive.NilObjectID, err
    }

    id := primitive.NewObjectID()
    inventory := models.Inventory{
        ID:           id,
        MotorcycleID: motorcycleID,
        Quantity:     input.Quantity,
        Location:     input.Location,
        LastUpdated:  time.Now().Format(time.RFC3339),
    }

    collection := database.DB.Collection("inventory")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err = collection.InsertOne(ctx, inventory)
    return id, err
}

// RemoveInventory deletes an inventory record from the database
func RemoveInventory(id string) error {
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }

    collection := database.DB.Collection("inventory")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err = collection.DeleteOne(ctx, bson.M{"_id": objectID})
    return err
}

// UpdateInventory modifies an inventory record in the database
func UpdateInventory(id string, input InventoryInput) error {
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }

    motorcycleID, err := primitive.ObjectIDFromHex(input.MotorcycleID)
    if err != nil {
        return err
    }

    collection := database.DB.Collection("inventory")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    update := bson.M{
        "$set": bson.M{
            "motorcycleId": motorcycleID,
            "quantity":     input.Quantity,
            "location":     input.Location,
            "lastUpdated":  time.Now().Format(time.RFC3339),
        },
    }

    _, err = collection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
    return err
}

// GetInventoryByID retrieves an inventory record by its ID
func GetInventoryByID(id string) (models.Inventory, error) {
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return models.Inventory{}, err
    }

    var inventory models.Inventory
    collection := database.DB.Collection("inventory")

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&inventory)
    return inventory, err
}