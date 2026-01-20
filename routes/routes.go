package routes

import (
	"github.com/gin-gonic/gin"
	"go-todo/controllers"
)



func SetupRoutes(r *gin.Engine) {
	// Web routes
	r.GET("/", controllers.Index)
	
	// API routes
	r.GET("/todos", controllers.GetTodos)
	r.POST("/todos", controllers.CreateTodo)
	r.PUT("/todos/:id", controllers.UpdateTodo)
	r.DELETE("/todos/:id", controllers.DeleteTodo)
}
