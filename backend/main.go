package main

import (
	"log"
	"os"

	"github.com/2SSK/portfolio_terminal/backend/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	//	 Load the environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create a new Fiber server
	server := fiber.New()
	PORT := os.Getenv("PORT")

	// Enable CORS for all origins
	server.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Setup the routes
	router.SetupRoutes(server)

	// Start the server
	server.Listen(":" + PORT)
}
