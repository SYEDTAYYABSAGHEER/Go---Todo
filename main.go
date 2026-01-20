package main

import (
	"github.com/gin-gonic/gin"
	"go-todo/database"
	"go-todo/routes"
)

func main() {
	r := gin.Default()

	// Load HTML templates
	r.LoadHTMLGlob("templates/*")
	
	// Serve static files (favicon)
	r.StaticFile("/favicon.ico", "./favicon.ico")

	database.ConnectDatabase()
	routes.SetupRoutes(r)

	r.Run(":8080")
}
