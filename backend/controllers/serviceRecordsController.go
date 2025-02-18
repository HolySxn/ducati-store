package controllers

import (
	"ducati-store/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetAllServiceRecords handles GET requests to retrieve all service records
func GetAllServiceRecords(c *gin.Context) {
	serviceRecords, err := services.GetAllServiceRecords()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch service records"})
		return
	}
	c.JSON(http.StatusOK, serviceRecords)
}

// CreateServiceRecord handles POST requests to add a new service record
func CreateServiceRecord(c *gin.Context) {
	var input services.ServiceRecordInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	id, err := services.CreateServiceRecord(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add service record"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Service record added successfully", "id": id})
}

// DeleteServiceRecord handles DELETE requests to remove a service record
func DeleteServiceRecord(c *gin.Context) {
	id := c.Param("id")
	if err := services.RemoveServiceRecord(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete service record"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Service record deleted successfully"})
}

// UpdateServiceRecord handles PUT requests to modify a service record
func UpdateServiceRecord(c *gin.Context) {
	id := c.Param("id")
	var input services.ServiceRecordInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := services.UpdateServiceRecord(id, input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update service record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Service record updated successfully"})
}

// GetServiceRecordByID handles GET requests to retrieve a service record by its ID
func GetServiceRecordByID(c *gin.Context) {
	id := c.Param("id")
	serviceRecord, err := services.GetServiceRecordByID(id)
	if err != nil {
		if err == primitive.ErrInvalidHex {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch service record"})
		}
		return
	}
	c.JSON(http.StatusOK, serviceRecord)
}
