package router

import (
	"github.com/Arshad-Siddiqui/go-bnb/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
)

func New() *fiber.App {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		AppName: "gobnb",
		Views:   engine,
	})
	app.Use(logger.New())
	app.Static("/", "./assets")
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})

	addUserRoutes(app)

	app.Post("/listings", func(c *fiber.Ctx) error {
		return controllers.ListingCreate(c)
	})

	app.Get("/listings", func(c *fiber.Ctx) error {
		return controllers.ListingIndex(c)
	})

	app.Delete("/listings/:id", func(c *fiber.Ctx) error {
		return controllers.ListingDelete(c)
	})
	return app
}
