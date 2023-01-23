package router

import (
	"github.com/Arshad-Siddiqui/go-bnb/controllers"
	"github.com/gofiber/fiber/v2"
)

func addListingRoutes(app *fiber.App) {
	listing := app.Group("listing")

	listing.Post("/create", controllers.ListingCreate)

	listing.Get("/index", controllers.ListingIndex)

	listing.Delete("/:id", controllers.ListingDelete)
}
