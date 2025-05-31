package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thanthtooaung-coding/go-crud-postgresql/controllers"
)

func TodoRoutes(r *gin.Engine) {
	todoGroup := r.Group("/todos")
	{
		todoGroup.POST("", controllers.CreateTodo)
		todoGroup.GET("", controllers.RetrieveAllTodoList)
		todoGroup.GET("/:id", controllers.RetrieveOneTodo)
		todoGroup.PUT("/:id", controllers.UpdateTodo)
		todoGroup.DELETE("/:id", controllers.DeleteTodo)
	}
}
