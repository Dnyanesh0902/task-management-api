package main

import (
	"log"
	"task-management-api/cache"
	"task-management-api/database"
	"task-management-api/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")

	}

	database.ConnectDB()
	cache.ConnectRedis()
	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run(":8080")
}