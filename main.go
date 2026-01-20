package main

import (
	"github.com/gin-gonic/gin"
	"go-todo/database"
	"go-todo/routes"
)

func main() {
	r := gin.Default()

	database.ConnectDatabase()
	routes.SetupRoutes(r)

	r.Run(":8080")
}
