package main

import (
	"log"

	"task-management-api/cache"
	"task-management-api/database"
	_ "task-management-api/docs"
	"task-management-api/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Task Management API
// @version 1.0
// @description Go Gin Task Management API with SQL Server, Redis, JWT, Docker, Repository Pattern
// @host localhost:8080
// @BasePath /api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}

	database.ConnectDB()
	cache.ConnectRedis()

	router := gin.Default()

	routes.SetupRoutes(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")
}