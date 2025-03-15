package config

import (
	"log"

	"github.com/2SSK/portfolio_terminal/backend/prisma/db"
)

var PrismaClient *db.PrismaClient

func InitDB() {
	PrismaClient = db.NewClient()
	if err := PrismaClient.Prisma.Connect(); err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
}

func CloseDB() {
	if PrismaClient != nil {
		PrismaClient.Prisma.Disconnect()
	}
}
