package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var db *gorm.DB

func connect (){
	err := godotenv.Load()

	 if err != nil {
        log.Fatal("Error loading .env file")
    }

	dsn := os.Getenv("DATABASE_URL")

	 database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
	
    log.Println("Connected to Postgres")
    db = database
}