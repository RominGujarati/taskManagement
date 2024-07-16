package database

import (
	"log"
	"taskManagement/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

// Connect establishes a connection to the PostgreSQL database
func Connect() {
	var err error
	dsn := "user=postgres password=123 dbname=task sslmode=disable"
	DB, err = gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	DB.AutoMigrate(&models.Task{})
}
