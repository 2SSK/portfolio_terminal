package db

import (
	"github.com/2SSK/portfolio_terminal/backend/models"
	"log"

	"gorm.io/driver/mongodb"
	"gorm.io/gorm"
)

var DB *gorm.DB
