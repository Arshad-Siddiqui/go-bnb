package main

import (
	"log"

	"github.com/Arshad-Siddiqui/go-bnb/initializers"
	"github.com/Arshad-Siddiqui/go-bnb/router"

	"os"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.InitDB()
}

func main() {
	app := router.New()

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
