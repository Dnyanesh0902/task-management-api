package database

import (
	"fmt"
	"log"
	"os"

	"task-management-api/models"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB * gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf(
		"sqlserver://%s:%s@%s:%s?database=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed To Connect Database :", err)
		
	}
	DB = db

	err = DB.AutoMigrate(&models.User{}, models.Project{}, models.Task{})
	
	if err != nil {
		log.Fatal("Migration Failed :", err)
	}
	fmt.Println("SQL Server Connected Successfully.")
}