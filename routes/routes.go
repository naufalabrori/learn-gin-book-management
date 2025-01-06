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

	bookGroup := r.Group("/books")
	{
		bookGroup.GET("/", controllers.GetBooks)
		bookGroup.POST("/", controllers.CreateBook)
		bookGroup.GET("/:id", controllers.GetBookByID)
		bookGroup.PUT("/:id", controllers.UpdateBook)
		bookGroup.DELETE("/:id", controllers.DeleteBook)
	}

	transactionGroup := r.Group("/transactions")
	{
		transactionGroup.GET("/", controllers.GetTransactions)
		transactionGroup.POST("/", controllers.CreateTransaction)
		transactionGroup.GET("/:id", controllers.GetTransactionByID)
		transactionGroup.PUT("/:id", controllers.UpdateTransaction)
		transactionGroup.DELETE("/:id", controllers.DeleteTransaction)
		transactionGroup.POST("/return/:id", controllers.ReturnTransaction)
	}

	return r
}
