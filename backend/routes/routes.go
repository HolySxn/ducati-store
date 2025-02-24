package routes

import (
	"ducati-store/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) *gin.Engine {
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
	router.GET("/inventory", controllers.GetAllInventory)
	router.POST("/inventory", controllers.CreateInventory)
	router.DELETE("/inventory/:id", controllers.DeleteInventory)
	router.PUT("/inventory/:id", controllers.UpdateInventory)
	router.GET("/inventory/:id", controllers.GetInventoryByID)

	// Customers
	router.GET("/customers", controllers.GetAllCustomers)
	router.POST("/customers", controllers.CreateCustomer)
	router.DELETE("/customers/:id", controllers.DeleteCustomer)
	router.PUT("/customers/:id", controllers.UpdateCustomer)
	router.GET("/customers/:id", controllers.GetCustomerByID)

	// Orders
	router.GET("/orders", controllers.GetAllOrders)
	router.POST("/orders", controllers.CreateOrder)
	router.DELETE("/orders/:id", controllers.DeleteOrder)
	router.PUT("/orders/:id", controllers.UpdateOrder)
	router.GET("/orders/:id", controllers.GetOrderByID)

	// Service Records
	router.GET("/service", controllers.GetAllServiceRecords)
	router.POST("/service", controllers.CreateServiceRecord)
	router.DELETE("/service/:id", controllers.DeleteServiceRecord)
	router.PUT("/service/:id", controllers.UpdateServiceRecord)
	router.GET("/service/:id", controllers.GetServiceRecordByID)

	return router
}
