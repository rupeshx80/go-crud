package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rupeshx80/go-crud/pkg/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default() // Creates a Gin router with default middleware (logger, recovery)

	// Group user-related routes
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/", controller.CreateUserController) // Create a new user
		// You can add more routes here, e.g.,
		// userRoutes.GET("/:id", controller.GetUserController)
		// userRoutes.PUT("/:id", controller.UpdateUserController)
		// userRoutes.DELETE("/:id", controller.DeleteUserController)
	}

	return r
}
