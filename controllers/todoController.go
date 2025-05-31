package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/thanthtooaung-coding/go-crud-postgresql/initializers"
	"github.com/thanthtooaung-coding/go-crud-postgresql/models"
	"net/http"
)

func CreateTodo(c *gin.Context) {
	var body struct {
		Content string `json:"content" binding:"required"`
		Status  *bool  `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if body.Status == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "status field is required but was not provided"})
		return
	}

	todo := models.Todo{Content: body.Content, Status: *body.Status}
	result := initializers.DB.Create(&todo)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "created",
		"todo":   todo,
	})
}

func RetrieveAllTodoList(c *gin.Context) {
	var todos []models.Todo
	initializers.DB.Find(&todos)

	c.JSON(http.StatusOK, gin.H{
		"status": "found",
		"todos":  todos,
	})
}

func RetrieveOneTodo(c *gin.Context) {
	id := c.Param("id")

	var todo models.Todo
	initializers.DB.First(&todo, id)

	c.JSON(http.StatusOK, gin.H{
		"status": "found",
		"todo":   todo,
	})
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Content string `json:"content" binding:"required"`
		Status  bool   `json:"status" binding:"required"`
	}
	err := c.BindJSON(&body)
	if err != nil {
		return
	}

	var todo models.Todo
	initializers.DB.First(&todo, id)

	initializers.DB.Model(&todo).Updates(models.Todo{
		Content: body.Content,
		Status:  body.Status,
	})

	c.JSON(http.StatusOK, gin.H{
		"status": "updated",
		"todo":   todo,
	})
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Todo{}, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Todo removed Successfully",
	})
}
