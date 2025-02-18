package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Motorcycle Model
type Motorcycle struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Model          string             `bson:"model"`
	Brand          string             `bson:"brand"`
	Year           int                `bson:"year"`
	Price          float64            `bson:"price"`
	Specifications []string           `bson:"specifications"`
	Status         string             `bson:"status"`
	CategoryID     primitive.ObjectID `bson:"categoryId"`
}

// Category Model
type Category struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
}

// Inventory Model
type Inventory struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	MotorcycleID primitive.ObjectID `bson:"motorcycleId"`
	Quantity     int                `bson:"quantity"`
	Location     string             `bson:"location"`
	LastUpdated  string             `bson:"lastUpdated"`
}

// Customer Model
type Customer struct {
	ID              primitive.ObjectID   `bson:"_id,omitempty"`
	Name            string               `bson:"name"`
	Email           string               `bson:"email"`
	Addresses       []string             `bson:"addresses"`
	PhoneNumbers    []string             `bson:"phoneNumbers"`
}

// Order Model
type Order struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	CustomerID primitive.ObjectID `bson:"customerId"`
	OrderDate  string             `bson:"orderDate"`
	Status     string             `bson:"status"`
	Items      []OrderItem        `bson:"items"`
	Payment    PaymentInfo        `bson:"payment"`
	Shipping   ShippingInfo       `bson:"shipping"`
}

type OrderItem struct {
	MotorcycleID primitive.ObjectID `bson:"motorcycleId"`
	Quantity     int                `bson:"quantity"`
}

type PaymentInfo struct {
	Method string  `bson:"method"`
	Amount float64 `bson:"amount"`
}

type ShippingInfo struct {
	Address string `bson:"address"`
	Status  string `bson:"status"`
}

// Service Record Model
type ServiceRecord struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	MotorcycleID primitive.ObjectID `bson:"motorcycleId"`
	CustomerID   primitive.ObjectID `bson:"customerId"`
	ServiceDate  string             `bson:"serviceDate"`
	Services     []string           `bson:"services"`
	Cost         float64            `bson:"cost"`
	Status       string             `bson:"status"`
}
