package router

import (
	"github.com/Arshad-Siddiqui/go-bnb/controllers"
	"github.com/gofiber/fiber/v2"
)

func addUserRoutes(app *fiber.App) {
	user := app.Group("user")

	user.Get("/signup", func(c *fiber.Ctx) error {
		return c.Render("signup", fiber.Map{})
	})

	user.Post("/signup", func(c *fiber.Ctx) error {
		return controllers.UserCreate(c)
	})

	user.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("login", fiber.Map{})
	})

	user.Post("/login", func(c *fiber.Ctx) error {
		return controllers.UserAuthenticate(c)
	})

	user.Get("/index", func(c *fiber.Ctx) error {
		return controllers.UserIndex(c)
	})

}
