package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "GoBnB",
	})

	app.Static("/signup", "./static/signup", fiber.Static{
		Index: "signup.html",
	})
	app.Static("/", "./static/home")
	log.Fatal(app.Listen(":3000"))
}
