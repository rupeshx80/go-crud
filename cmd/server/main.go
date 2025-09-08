package main

import (
    "github.com/rupeshx80/go-crud/pkg/db"
    "github.com/rupeshx80/go-crud/pkg/models"
    "log"

    "github.com/gin-gonic/gin"
)

func main() {
    // Connect to Postgres
    db.Connect()

    err := db.DB.AutoMigrate(&models.User{})
    if err != nil {
        log.Fatal("Migration failed:", err)
    }
    log.Println("Database migrated")

    // Setup Gin
    r := gin.Default()

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })

    r.Run(":4000")
}
