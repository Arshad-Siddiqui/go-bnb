package main

import (
	"fmt"

	"github.com/Arshad-Siddiqui/go-bnb/initializers"
	"github.com/Arshad-Siddiqui/go-bnb/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.InitDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.Listing{})
	fmt.Println("Migrated")
}
