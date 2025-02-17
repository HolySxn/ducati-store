package controllers

import (
	"ducati-store/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetAllInventory handles GET requests to retrieve all inventory records
func GetAllInventory(c *gin.Context) {
	inventory, err := services.GetAllInventory()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch inventory"})
		return
	}
	c.JSON(http.StatusOK, inventory)
}

// CreateInventory handles POST requests to add a new inventory record
func CreateInventory(c *gin.Context) {
	var input services.InventoryInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	id, err := services.CreateInventory(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add inventory"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Inventory added successfully", "id": id})
}

// DeleteInventory handles DELETE requests to remove an inventory record
func DeleteInventory(c *gin.Context) {
	id := c.Param("id")
	if err := services.RemoveInventory(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete inventory"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Inventory deleted successfully"})
}

// UpdateInventory handles PUT requests to modify an inventory record
func UpdateInventory(c *gin.Context) {
	id := c.Param("id")
	var input services.InventoryInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := services.UpdateInventory(id, input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update inventory"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inventory updated successfully"})
}

// GetInventoryByID handles GET requests to retrieve an inventory record by its ID
func GetInventoryByID(c *gin.Context) {
	id := c.Param("id")
	inventory, err := services.GetInventoryByID(id)
	if err != nil {
		if err == primitive.ErrInvalidHex {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch inventory"})
		}
		return
	}
	c.JSON(http.StatusOK, inventory)
}
