package main

import (
	"github.com/thanthtooaung-coding/go-crud-postgresql/initializers"
	"github.com/thanthtooaung-coding/go-crud-postgresql/models"
	"log"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	err := initializers.DB.AutoMigrate(&models.Todo{})
	if err != nil {
		log.Fatalf("❌ Failed to migrate: %v", err)
	}
	log.Println("✅ Migration completed successfully.")
}
