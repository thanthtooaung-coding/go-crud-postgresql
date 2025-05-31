package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thanthtooaung-coding/go-crud-postgresql/initializers"
	"github.com/thanthtooaung-coding/go-crud-postgresql/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()
	routes.TodoRoutes(r)

	err := r.Run()
	if err != nil {
		return
	}
}
