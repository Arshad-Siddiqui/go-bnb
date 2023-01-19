package router

import (
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
	addListingRoutes(app)

	return app
}
