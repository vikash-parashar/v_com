package config

import (
	"fmt"
	"log"
	"os"
	"v_com/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// Database connection parameters
	dbHost := "localhost"
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	// Construct the database connection URL
	dbURL := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		dbHost, dbPort, dbName, dbUser, dbPassword)
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)

	}
	DB = db
	// Migrate your database tables here if needed
	db.AutoMigrate(&models.Address{}, &models.CartItem{}, &models.Order{}, &models.User{}, &models.Product{}, &models.OrderProduct{})

}
