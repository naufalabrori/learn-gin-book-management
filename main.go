package main

import (
	"learn-go-gin/config"
	"learn-go-gin/database"
	"learn-go-gin/routes"
	"time"

	"github.com/gin-contrib/cors"
)

func main() {
	config.ConnectDB()
	database.Migrate()
	r := routes.SetupRoutes()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},        // Origin frontend Anda
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"}, // Metode HTTP yang diizinkan
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour, // Cache preflight request selama 12 jam
	}))

	r.Run(":8080")
}
