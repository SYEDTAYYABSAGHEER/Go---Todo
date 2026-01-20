package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-todo/database"
	"go-todo/models"
)

// GET / - Render home page with todos
func Index(c *gin.Context) {
	var todos []models.Todo
	database.DB.Order("created_at desc").Find(&todos)
	c.HTML(http.StatusOK, "base.html", gin.H{
		"Title": "Home",
		"Todos": todos,
	})
}

// GET /todos - API endpoint
func GetTodos(c *gin.Context) {
	var todos []models.Todo
	database.DB.Find(&todos)
	c.JSON(http.StatusOK, todos)
}

// POST /todos
func CreateTodo(c *gin.Context) {
	var todo models.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&todo)
	c.JSON(http.StatusCreated, todo)
}

// PUT /todos/:id
func UpdateTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")

	if err := database.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	c.ShouldBindJSON(&todo)
	database.DB.Save(&todo)

	c.JSON(http.StatusOK, todo)
}

// DELETE /todos/:id
func DeleteTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")

	if err := database.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	database.DB.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}
