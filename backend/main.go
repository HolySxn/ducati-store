package main

import (
	"ducati-store/database"
	"ducati-store/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	

	database.ConnectDB()

	router := gin.Default()

	router.Static("/static", "./frontend/build/static")
	router.StaticFile("/", "./frontend/build/index.html")
	
	routes.SetupRouter(router)
	router.Run(":8080")
}
