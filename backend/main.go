package main

import (
	"ducati-store/database"
	"ducati-store/routes"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectDB()

	router := gin.Default()

	// CORS Middleware
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // React Frontend URL
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Use CORS Middleware
	router.Use(func(c *gin.Context) {
		corsMiddleware.HandlerFunc(c.Writer, c.Request)
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})

	routes.SetupRouter(router)
	router.Run(":8080")
}
