package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rupeshx80/go-crud/pkg/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default() 

	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/", controller.CreateUserController)
		userRoutes.GET("/email", controller.GetUserByEmailController) 
		userRoutes.GET("/username", controller.GetUserByUsernameController) 
		userRoutes.PUT("/:id", controller.UpdateUserController)

	}

	return r
}
