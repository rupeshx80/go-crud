package main

import (
    "github.com/rupeshx80/go-crud/pkg/db"
    "github.com/rupeshx80/go-crud/pkg/models"
    "log"

    "github.com/rupeshx80/go-crud/pkg/router"
)

func main() {
    db.Connect()

    err := db.RJ.AutoMigrate(&models.User{})
    if err != nil {
        log.Fatal("Migration failed:", err)
    }
    log.Println("Database migrated")

   r := router.SetupRouter()

    r.Run(":4040")
}
