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

	user.Post("/signup", controllers.UserCreate)

	user.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("login", fiber.Map{})
	})

	user.Post("/login", controllers.UserAuthenticate)

	user.Get("/index", controllers.UserIndex)

}
