package router

import (
	"github.com/Arshad-Siddiqui/go-bnb/controllers"
	"github.com/gofiber/fiber/v2"
)

func addListingRoutes(app *fiber.App) {
	listing := app.Group("listing")

	listing.Post("/create", func(c *fiber.Ctx) error {
		return controllers.ListingCreate(c)
	})

	listing.Get("/index", func(c *fiber.Ctx) error {
		return controllers.ListingIndex(c)
	})

	listing.Delete("/:id", func(c *fiber.Ctx) error {
		return controllers.ListingDelete(c)
	})
}
