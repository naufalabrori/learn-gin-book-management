package main

import (
	"learn-go-gin/config"
	"learn-go-gin/database"
	"learn-go-gin/routes"
)

func main() {
	config.ConnectDB()
	database.Migrate()
	r := routes.SetupRoutes()
	r.Run(":8080")
}
