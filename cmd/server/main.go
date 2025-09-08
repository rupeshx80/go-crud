package main

import (
    "github.com/rupeshx80/go-crud/pkg/db"
    "github.com/rupeshx80/go-crud/pkg/models"
    "github.com/rupeshx80/go-crud/pkg/controller"
    "log"

    "github.com/gin-gonic/gin"
)

func main() {
    db.Connect()

    err := db.RJ.AutoMigrate(&models.User{})
    if err != nil {
        log.Fatal("Migration failed:", err)
    }
    log.Println("Database migrated")

    r := gin.Default()

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })

    userRoutes := r.Group("/users")
	{
		userRoutes.POST("/", controller.CreateUserController)
		//add more routes herefor GET, PUT, DELETE,etc
	}

    r.Run(":4040")
}
