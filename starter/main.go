// main.go
package main

import (
	"cms-backend/models"
	"cms-backend/routes"
	"cms-backend/utils"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

// @title CMS Backend API
// @version 1.0
// @description This is a backend API for a Content Management System (CMS).
// @host localhost:8080
// @BasePath /api/v1

func main() {
	// Initialize database connection
	db, err := utils.ConnectDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Get the underlying *sql.DB instance and defer its closure
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}
	defer sqlDB.Close()

	// Get the environment variable
	env := os.Getenv("ENV")
	if env == "" {
		env = "development" // default to development if ENV is not set
	}

	// Conditionally run AutoMigrate in development environment
	if env == "development" {
		log.Println("Running AutoMigrate...")
		if err := db.AutoMigrate(&models.Page{}, &models.Post{}, &models.Media{}); err != nil {
			log.Fatalf("Failed to automigrate database: %v", err)
		}
	}

	// Set Gin mode based on environment
	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	// Initialize routes
	routes.InitializeRoutes(router, db)

	// Run the server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}