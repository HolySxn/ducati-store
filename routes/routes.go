package routes

import (
	"ducati-store/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/motorcycles", controllers.GetMotorcycles)
	router.POST("/motorcycles", controllers.AddMotorcycle)
	router.DELETE("/motorcycles/:id", controllers.DeleteMotorcycle)
	router.PUT("/motorcycles/:id", controllers.UpdateMotorcycle)
	router.GET("/motorcycles/:id", controllers.GetMotorcycleByID)

	return router
}
