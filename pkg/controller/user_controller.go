package controller

import (
	"net/http"

	"strconv"
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

func GetUserByEmailController(c *gin.Context) {
	// Get "email" from query params -> /users/email?email=test@gmail.com
	email := c.Query("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email query parameter is required"})
		return
	}

	user, err := service.GetUserByEmailService(email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func GetUserByUsernameController (c *gin.Context){
	username := c.Query(("username"))

	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username query parameter is required"})
		return
	}

	user,err := service.GetUserByUsernameService((username))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})

}

func UpdateUserController(c *gin.Context) {
	// Get ID from URL param
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	// Parse JSON body
	var newData map[string]interface{}
	if err := c.ShouldBindJSON(&newData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	// Call service
	updatedUser, err := service.UpdateUserService(uint(id), newData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Success response
	c.JSON(http.StatusOK, gin.H{"user": updatedUser})
}