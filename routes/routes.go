package routes

import (
	"learn-go-gin/controllers"
	"learn-go-gin/utils"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) *gin.Engine {
	// r := gin.Default()

	publicGroup := r.Group("/")
	{
		publicGroup.POST("/users/login", controllers.Login)
		publicGroup.POST("/users", controllers.CreateUser)
		publicGroup.GET("/categories", controllers.GetCategories)
		publicGroup.GET("/categories/:id", controllers.GetCategoryByID)
		publicGroup.GET("/books", controllers.GetBooks)
		publicGroup.GET("/books/:id", controllers.GetBookByID)
	}

	userGroup := r.Group("/users").Use(utils.JWTAuthMiddleware())
	{
		userGroup.GET("/", controllers.GetUsers)
		userGroup.GET("/:id", controllers.GetUserByID)
		userGroup.PUT("/:id", controllers.UpdateUser)
		userGroup.DELETE("/:id", controllers.DeleteUser)
		userGroup.POST("/images/:id", controllers.UploadUserImage)
		userGroup.POST("/change-password/:id", controllers.ChangePassword)
	}

	categoryGroup := r.Group("/categories").Use(utils.JWTAuthMiddleware())
	{
		categoryGroup.POST("/", controllers.CreateCategory)
		categoryGroup.PUT("/:id", controllers.UpdateCategory)
		categoryGroup.DELETE("/:id", controllers.DeleteCategory)
	}

	bookGroup := r.Group("/books").Use(utils.JWTAuthMiddleware())
	{
		bookGroup.POST("/", controllers.CreateBook)
		bookGroup.PUT("/:id", controllers.UpdateBook)
		bookGroup.DELETE("/:id", controllers.DeleteBook)
	}

	transactionGroup := r.Group("/transactions").Use(utils.JWTAuthMiddleware())
	{
		transactionGroup.GET("/", controllers.GetTransactions)
		transactionGroup.POST("/", controllers.CreateTransaction)
		transactionGroup.GET("/:id", controllers.GetTransactionByID)
		transactionGroup.PUT("/:id", controllers.UpdateTransaction)
		transactionGroup.DELETE("/:id", controllers.DeleteTransaction)
		transactionGroup.POST("/return/:id", controllers.ReturnTransaction)
	}

	finesGroup := r.Group("/fines").Use(utils.JWTAuthMiddleware())
	{
		finesGroup.GET("/", controllers.GetFines)
		finesGroup.POST("/", controllers.CreateFines)
		finesGroup.GET("/:id", controllers.GetFinesByID)
		finesGroup.PUT("/:id", controllers.UpdateFines)
		finesGroup.DELETE("/:id", controllers.DeleteFines)
		finesGroup.GET("/transaction/:transactionId", controllers.GetFinesByTransactionID)
	}

	return r
}
