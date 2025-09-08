package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rupeshx80/go-crud/pkg/models"
	"github.com/rupeshx80/go-crud/pkg/service"
)

func CreateUserController(c *gin.Context) {
	var user models.User

	//binds json-body into go structs
	err := c.ShouldBindJSON(&user) 

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return 
	}

	createdUser, err := service.CreateUserService(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": createdUser})
}
