package controllers

import (
	"ducati-store/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get all motorcycles
func GetMotorcycles(c *gin.Context) {
	motorcycles, err := services.GetAllMotorcycles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch motorcycles"})
		return
	}
	c.JSON(http.StatusOK, motorcycles)
}

// Add a new motorcycle
func AddMotorcycle(c *gin.Context) {
	var motorcycle services.MotorcycleInput
	if err := c.BindJSON(&motorcycle); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	id, err := services.CreateMotorcycle(motorcycle)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add motorcycle"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Motorcycle added successfully", "id": id})
}

// Delete a motorcycle
func DeleteMotorcycle(c *gin.Context) {
	id := c.Param("id")
	if err := services.RemoveMotorcycle(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete motorcycle"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Motorcycle deleted successfully"})
}

// Update a motorcycle
func UpdateMotorcycle(c *gin.Context) {
	id := c.Param("id")
	var motorcycle services.MotorcycleInput
	if err := c.BindJSON(&motorcycle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := services.UpdateMotorcycle(id, motorcycle); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update motorcycle"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Motorcycle updated successfully"})
}

// GetMotorcycleByID retrieves a motorcycle by its ID
func GetMotorcycleByID(c *gin.Context) {
	id := c.Param("id")
	motorcycle, err := services.GetMotorcycleByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Motorcycle not found"})
		return
	}
	c.JSON(http.StatusOK, motorcycle)
}
