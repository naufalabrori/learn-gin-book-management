package routes

import (
	"learn-go-gin/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	userGroup := r.Group("/users")
	{
		userGroup.GET("/", controllers.GetUsers)
		userGroup.POST("/", controllers.CreateUser)
		userGroup.GET("/:id", controllers.GetUserByID)
		userGroup.PUT("/:id", controllers.UpdateUser)
		userGroup.DELETE("/:id", controllers.DeleteUser)
		userGroup.POST("/login", controllers.Login)
	}

	categoryGroup := r.Group("/categories")
	{
		categoryGroup.GET("/", controllers.GetCategories)
		categoryGroup.POST("/", controllers.CreateCategory)
		categoryGroup.GET("/:id", controllers.GetCategoryByID)
		categoryGroup.PUT("/:id", controllers.UpdateCategory)
		categoryGroup.DELETE("/:id", controllers.DeleteCategory)
	}

	return r
}
