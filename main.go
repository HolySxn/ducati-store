package main

import (
	"ducati-store/database"
	"ducati-store/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectDB()
	router := routes.SetupRouter()
	router.Run(":8080")
}
