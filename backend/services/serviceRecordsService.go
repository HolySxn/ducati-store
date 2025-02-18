package services

import (
    "context"
    "ducati-store/database"
    "ducati-store/models"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

// ServiceRecordInput represents the input data for creating or updating a service record
type ServiceRecordInput struct {
    MotorcycleID string   `json:"motorcycleId"`
    CustomerID   string   `json:"customerId"`
    ServiceDate  string   `json:"serviceDate"`
    Services     []string `json:"services"`
    Cost         float64  `json:"cost"`
    Status       string   `json:"status"`
}

// GetAllServiceRecords retrieves all service records from the database
func GetAllServiceRecords() ([]models.ServiceRecord, error) {
    var serviceRecords []models.ServiceRecord
    collection := database.DB.Collection("serviceRecords")

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    for cursor.Next(ctx) {
        var serviceRecord models.ServiceRecord
        cursor.Decode(&serviceRecord)
        serviceRecords = append(serviceRecords, serviceRecord)
    }

    return serviceRecords, nil
}

// CreateServiceRecord adds a new service record to the database
func CreateServiceRecord(input ServiceRecordInput) (primitive.ObjectID, error) {
    motorcycleID, err := primitive.ObjectIDFromHex(input.MotorcycleID)
    if err != nil {
        return primitive.NilObjectID, err
    }

    customerID, err := primitive.ObjectIDFromHex(input.CustomerID)
    if err != nil {
        return primitive.NilObjectID, err
    }

    id := primitive.NewObjectID()
    serviceRecord := models.ServiceRecord{
        ID:           id,
        MotorcycleID: motorcycleID,
        CustomerID:   customerID,
        ServiceDate:  input.ServiceDate,
        Services:     input.Services,
        Cost:         input.Cost,
        Status:       input.Status,
    }

    collection := database.DB.Collection("serviceRecords")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err = collection.InsertOne(ctx, serviceRecord)
    return id, err
}

// RemoveServiceRecord deletes a service record from the database
func RemoveServiceRecord(id string) error {
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }

    collection := database.DB.Collection("serviceRecords")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err = collection.DeleteOne(ctx, bson.M{"_id": objectID})
    return err
}

// UpdateServiceRecord modifies a service record in the database
func UpdateServiceRecord(id string, input ServiceRecordInput) error {
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }

    motorcycleID, err := primitive.ObjectIDFromHex(input.MotorcycleID)
    if err != nil {
        return err
    }

    customerID, err := primitive.ObjectIDFromHex(input.CustomerID)
    if err != nil {
        return err
    }

    collection := database.DB.Collection("serviceRecords")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    update := bson.M{
        "$set": bson.M{
            "motorcycleId": motorcycleID,
            "customerId":   customerID,
            "serviceDate":  input.ServiceDate,
            "services":     input.Services,
            "cost":         input.Cost,
            "status":       input.Status,
        },
    }

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
    return err
}

// GetServiceRecordByID retrieves a service record by its ID
func GetServiceRecordByID(id string) (models.ServiceRecord, error) {
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return models.ServiceRecord{}, err
    }

    var serviceRecord models.ServiceRecord
    collection := database.DB.Collection("serviceRecords")

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&serviceRecord)
    return serviceRecord, err
}