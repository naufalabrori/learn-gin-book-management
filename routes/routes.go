package routes

import (
	"learn-go-gin/controllers"
	"learn-go-gin/utils"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.Use(CORS())

	//Middleware CORS
	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"http://localhost:3000"},
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowHeaders:     []string{"Content-Type", "Authorization"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour,
	// }))

	// Handle OPTIONS (Preflight)
	// r.OPTIONS("/*path", func(c *gin.Context) {
	// 	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	// 	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	// 	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
	// 	c.Header("Access-Control-Allow-Credentials", "true")
	// 	c.Status(http.StatusNoContent)
	// })

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

	r.Static("/uploads", "./uploads")

	return r
}
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
