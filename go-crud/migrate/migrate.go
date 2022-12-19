package main

import (
	"github.com/ghana7989-go-crud/initializers"
	"github.com/ghana7989-go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
