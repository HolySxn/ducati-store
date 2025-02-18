package controllers

import (
	"ducati-store/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetAllCustomers handles GET requests to retrieve all customers
func GetAllCustomers(c *gin.Context) {
	customers, err := services.GetAllCustomers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch customers"})
		return
	}
	c.JSON(http.StatusOK, customers)
}

// CreateCustomer handles POST requests to add a new customer
func CreateCustomer(c *gin.Context) {
	var input services.CustomerInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	id, err := services.CreateCustomer(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add customer"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Customer added successfully", "id": id})
}

// DeleteCustomer handles DELETE requests to remove a customer
func DeleteCustomer(c *gin.Context) {
	id := c.Param("id")
	if err := services.RemoveCustomer(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete customer"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Customer deleted successfully"})
}

// UpdateCustomer handles PUT requests to modify a customer
func UpdateCustomer(c *gin.Context) {
	id := c.Param("id")
	var input services.CustomerInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := services.UpdateCustomer(id, input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update customer"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer updated successfully"})
}

// GetCustomerByID handles GET requests to retrieve a customer by its ID
func GetCustomerByID(c *gin.Context) {
	id := c.Param("id")
	customer, err := services.GetCustomerByID(id)
	if err != nil {
		if err == primitive.ErrInvalidHex {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch customer"})
		}
		return
	}
	c.JSON(http.StatusOK, customer)
}
