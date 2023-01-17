package controllers

import (
	"github.com/Arshad-Siddiqui/go-bnb/initializers"
	"github.com/Arshad-Siddiqui/go-bnb/models"
	"github.com/gofiber/fiber/v2"
)

// Signup function
func UserCreate(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return err
	}
	initializers.DB.Create(&user)
	return c.JSON(user)
}

func UserIndex(c *fiber.Ctx) error {
	var users []models.User
	initializers.DB.Find(&users)
	return c.JSON(users)
}
