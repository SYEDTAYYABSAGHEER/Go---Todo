package routes

import (
	"github.com/gin-gonic/gin"
	"go-todo/controllers"
)



func SetupRoutes(r *gin.Engine) {
	// Root route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to Go Todo API!"})
	})
	r.GET("/todos", controllers.GetTodos)
	r.POST("/todos", controllers.CreateTodo)
	r.PUT("/todos/:id", controllers.UpdateTodo)
	r.DELETE("/todos/:id", controllers.DeleteTodo)
}
