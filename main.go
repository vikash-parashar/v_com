package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// Load environment variables from a .env file in the current directory
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file:", err)
		os.Exit(1)
	}
}

func main() {

}
