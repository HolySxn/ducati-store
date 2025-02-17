package routes

import (
	"ducati-store/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Motocycles
	router.GET("/motorcycles", controllers.GetMotorcycles)
	router.POST("/motorcycles", controllers.AddMotorcycle)
	router.DELETE("/motorcycles/:id", controllers.DeleteMotorcycle)
	router.PUT("/motorcycles/:id", controllers.UpdateMotorcycle)
	router.GET("/motorcycles/:id", controllers.GetMotorcycleByID)

	// Categories
	router.GET("/categories", controllers.GetAllCategories)
	router.POST("/categories", controllers.CreateCategory)
	router.DELETE("/categories/:id", controllers.DeleteCategory)
	router.PUT("/categories/:id", controllers.UpdateCategory)
	router.GET("/categories/:id", controllers.GetCategoryByID)

	// Inventory
	

	return router
}
