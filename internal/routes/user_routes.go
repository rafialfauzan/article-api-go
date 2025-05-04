package routes

import (
	"golang-api/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(api *gin.RouterGroup, userHandler *handler.UserHandler) {
	users := api.Group("/users")
	{
		users.POST("", userHandler.CreateUser)
		users.GET("", userHandler.GetAllUsers)
		users.GET("/:id", userHandler.GetUserByID)
		users.PUT("/:id", userHandler.UpdateUser)
		users.DELETE("/:id", userHandler.DeleteUser)
	}
}
