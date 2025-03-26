package main

import (
	"log"
	"os"

	"github.com/2SSK/portfolio_terminal/backend/config"
	"github.com/2SSK/portfolio_terminal/backend/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Load the environment variables from .env file if it exists (optional)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}

	// Initialize database connection
	config.InitDB()
	defer config.CloseDB()

	// Create a new Fiber server
	server := fiber.New()

	// Enable CORS for all origins
	server.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Setup the routes
	router.SetupRoutes(server)

	// Start the server
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3000"
	}
	server.Listen(":" + PORT)
}
