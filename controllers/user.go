package controllers

import (
	"fmt"

	"github.com/Arshad-Siddiqui/go-bnb/initializers"
	"github.com/Arshad-Siddiqui/go-bnb/models"
	"github.com/gofiber/fiber/v2"
)

// Signup function
func UserCreate(c *fiber.Ctx) error {
	fmt.Println(string(c.Body()), "body")
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

func UserAuthenticate(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return err
	}
	var dbUser models.User
	initializers.DB.Where("email = ?", user.Email).First(&dbUser)
	if dbUser.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}
	if user.Password != dbUser.Password {
		return c.Status(401).JSON(fiber.Map{
			"message": "Incorrect password",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Logged in",
	})
}
