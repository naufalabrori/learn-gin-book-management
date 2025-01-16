package routes

import (
	"learn-go-gin/controllers"
	"learn-go-gin/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	//Middleware CORS
	r.Use(CORS())

	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"http://localhost:3000"},
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowHeaders:     []string{"Content-Type", "Authorization"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour,
	// }))

	// Public Routes
	r.POST("/users/login", controllers.Login)
	r.POST("/users", controllers.CreateUser)
	r.GET("/categories", controllers.GetCategories)
	r.GET("/categories/:id", controllers.GetCategoryByID)
	r.GET("/books", controllers.GetBooks)
	r.GET("/books/:id", controllers.GetBookByID)

	// User Routes (with JWT middleware)
	r.GET("/users", middlewares.JWTAuthMiddleware(), controllers.GetUsers)
	r.GET("/users/:id", middlewares.JWTAuthMiddleware(), controllers.GetUserByID)
	r.PUT("/users/:id", middlewares.JWTAuthMiddleware(), controllers.UpdateUser)
	r.DELETE("/users/:id", middlewares.JWTAuthMiddleware(), controllers.DeleteUser)
	r.POST("/users/images/:id", middlewares.JWTAuthMiddleware(), controllers.UploadUserImage)
	r.POST("/users/change-password/:id", middlewares.JWTAuthMiddleware(), controllers.ChangePassword)

	// Category Routes (with JWT middleware)
	r.POST("/categories", middlewares.JWTAuthMiddleware(), controllers.CreateCategory)
	r.PUT("/categories/:id", middlewares.JWTAuthMiddleware(), controllers.UpdateCategory)
	r.DELETE("/categories/:id", middlewares.JWTAuthMiddleware(), controllers.DeleteCategory)

	// Book Routes (with JWT middleware)
	r.POST("/books", middlewares.JWTAuthMiddleware(), controllers.CreateBook)
	r.PUT("/books/:id", middlewares.JWTAuthMiddleware(), controllers.UpdateBook)
	r.DELETE("/books/:id", middlewares.JWTAuthMiddleware(), controllers.DeleteBook)

	// Transaction Routes (with JWT middleware)
	r.GET("/transactions", middlewares.JWTAuthMiddleware(), controllers.GetTransactions)
	r.POST("/transactions", middlewares.JWTAuthMiddleware(), controllers.CreateTransaction)
	r.GET("/transactions/:id", middlewares.JWTAuthMiddleware(), controllers.GetTransactionByID)
	r.PUT("/transactions/:id", middlewares.JWTAuthMiddleware(), controllers.UpdateTransaction)
	r.DELETE("/transactions/:id", middlewares.JWTAuthMiddleware(), controllers.DeleteTransaction)
	r.POST("/transactions/return/:id", middlewares.JWTAuthMiddleware(), controllers.ReturnTransaction)

	// Fines Routes (with JWT middleware)
	r.GET("/fines", middlewares.JWTAuthMiddleware(), controllers.GetFines)
	r.POST("/fines", middlewares.JWTAuthMiddleware(), controllers.CreateFines)
	r.GET("/fines/:id", middlewares.JWTAuthMiddleware(), controllers.GetFinesByID)
	r.PUT("/fines/:id", middlewares.JWTAuthMiddleware(), controllers.UpdateFines)
	r.DELETE("/fines/:id", middlewares.JWTAuthMiddleware(), controllers.DeleteFines)
	r.GET("/fines/transaction/:transactionId", middlewares.JWTAuthMiddleware(), controllers.GetFinesByTransactionID)

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
