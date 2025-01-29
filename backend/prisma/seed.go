package main

import (
	"context"
	"fmt"
	"log"
	"os"

	db "github.com/2SSK/portfolio_terminal/backend/prisma/db"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// Create a prisma client
	client := db.NewClient()

	// connect to the database
	if err := client.Prisma.Connect(); err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer client.Prisma.Disconnect()

	ctx := context.Background()

	// Hardcoded admin credentials
	username := os.Getenv("ADMIN_USERNAME")
	password := os.Getenv("ADMIN_PASSWORD")

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed hash the password: %v", err)
	}

	// Upsert the admin user
	admin, err := client.Admin.UpsertOne(
		db.Admin.Username.Equals(username), // Check if an admin with the username exists
	).Update(
		db.Admin.Password.Set(string(hashedPassword)),
	).Create(
		db.Admin.Username.Set(username),
		db.Admin.Password.Set(string(hashedPassword)),
	).Exec(ctx)

	if err != nil {
		log.Fatalf("Failed to seed admin user: %v", err)
	}

	fmt.Printf("Admin user seeded: %v\n", admin)
}
