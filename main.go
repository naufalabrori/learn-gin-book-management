package main

import (
	"learn-go-gin/config"
	"learn-go-gin/database"
	"learn-go-gin/routes"
)

func main() {
	config.ConnectDB()
	database.Migrate()

	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"http://localhost:3000"},
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "Cache-Control", "Origin"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour,
	// }))

	r := routes.SetupRoutes()

	r.Run(":8080")
}
