package controllers

import (
	"github.com/Arshad-Siddiqui/go-bnb/initializers"
	"github.com/Arshad-Siddiqui/go-bnb/models"
	"github.com/gofiber/fiber/v2"
)

// ListingCreate function
func ListingCreate(c *fiber.Ctx) error {
	listing := new(models.Listing)
	if err := c.BodyParser(listing); err != nil {
		return err
	}
	initializers.DB.Create(&listing)
	return c.JSON(listing)
}

func ListingIndex(c *fiber.Ctx) error {
	var listings []models.Listing
	initializers.DB.Find(&listings)
	return c.JSON(listings)
}
