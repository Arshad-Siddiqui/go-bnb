package main

import (
	"log"

	"github.com/gofiber/template/html"

	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		AppName: "gobnb",
		Views:   engine,
	})

	app.Static("/", "./assets")
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})

	app.Get("/signup", func(c *fiber.Ctx) error {
		return c.Render("signup", fiber.Map{})
	})

	app.Post("/signup", func(c *fiber.Ctx) error {
		username := c.FormValue("username")
		password := c.FormValue("password")
		return c.SendString("Username: " + username + " Password: " + password)
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("login", fiber.Map{})
	})

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
