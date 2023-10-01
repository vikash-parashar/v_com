package main

import (
	"fmt"
	"os"
	"v_com/config"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// Load environment variables from a .env file in the current directory
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file:", err)
		os.Exit(1)
	}
	config.InitDB()
}

func main() {
	// Create a Gin router
	router := gin.Default()

	// Run the Gin server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	serverAddress := fmt.Sprintf(":%s", port)
	router.Run(serverAddress)
}
