package controllers

import (
	"ducati-store/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetAllCategories handles GET requests to retrieve all categories
func GetAllCategories(c *gin.Context) {
	categories, err := services.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories"})
		return
	}
	c.JSON(http.StatusOK, categories)
}

// CreateCategory handles POST requests to add a new category
func CreateCategory(c *gin.Context) {
	var input services.CategoryInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	id, err := services.CreateCategory(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add category"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Category added successfully", "id": id})
}

// DeleteCategory handles DELETE requests to remove a category
func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	if err := services.RemoveCategory(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}

// UpdateCategory handles PUT requests to modify a category
func UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var input services.CategoryInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := services.UpdateCategory(id, input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category updated successfully"})
}

// GetCategoryByID handles GET requests to retrieve a category by its ID
func GetCategoryByID(c *gin.Context) {
	id := c.Param("id")
	category, err := services.GetCategoryByID(id)
	if err != nil {
		if err == primitive.ErrInvalidHex {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch category"})
		}
		return
	}
	c.JSON(http.StatusOK, category)
}
