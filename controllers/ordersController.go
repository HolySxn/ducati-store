package controllers

import (
	"ducati-store/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetAllOrders handles GET requests to retrieve all orders
func GetAllOrders(c *gin.Context) {
	orders, err := services.GetAllOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}
	c.JSON(http.StatusOK, orders)
}

// CreateOrder handles POST requests to add a new order
func CreateOrder(c *gin.Context) {
	var input services.OrderInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	id, err := services.CreateOrder(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add order"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Order added successfully", "id": id})
}

// DeleteOrder handles DELETE requests to remove an order
func DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	if err := services.RemoveOrder(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}

// UpdateOrder handles PUT requests to modify an order
func UpdateOrder(c *gin.Context) {
	id := c.Param("id")
	var input services.OrderInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := services.UpdateOrder(id, input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order updated successfully"})
}

// GetOrderByID handles GET requests to retrieve an order by its ID
func GetOrderByID(c *gin.Context) {
	id := c.Param("id")
	order, err := services.GetOrderByID(id)
	if err != nil {
		if err == primitive.ErrInvalidHex {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch order"})
		}
		return
	}
	c.JSON(http.StatusOK, order)
}
