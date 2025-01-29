package main

import (
	"log"
	"os"

	"github.com/2SSK/portfolio_terminal/backend/router"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load the environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create a new Fiber app
	app := fiber.New()
	PORT := os.Getenv("PORT")

	// Setup the routes
	router.SetupRoutes(app)

	app.Listen(":" + PORT)
}
