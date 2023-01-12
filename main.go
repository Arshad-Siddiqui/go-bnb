package main

import (
	"log"

	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := fiber.New(fiber.Config{
		AppName: "GoBnB",
	})

	app.Static("/", "static")
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
