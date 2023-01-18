package main

import (
	"log"

	"github.com/Arshad-Siddiqui/go-bnb/controllers"
	"github.com/Arshad-Siddiqui/go-bnb/initializers"
	"github.com/gofiber/template/html"

	"os"

	"github.com/gofiber/fiber/v2"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.InitDB()
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
		return controllers.UserCreate(c)
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("login", fiber.Map{})
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		return controllers.UserAuthenticate(c)
	})

	app.Get("/users", func(c *fiber.Ctx) error {
		return controllers.UserIndex(c)
	})

	app.Post("/listings", func(c *fiber.Ctx) error {
		return controllers.ListingCreate(c)
	})

	app.Get("/listings", func(c *fiber.Ctx) error {
		return controllers.ListingIndex(c)
	})

	app.Delete("/listings/:id", func(c *fiber.Ctx) error {
		return controllers.ListingDelete(c)
	})
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
