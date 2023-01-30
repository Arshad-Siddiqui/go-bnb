package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/Arshad-Siddiqui/go-bnb/controllers"
	"github.com/Arshad-Siddiqui/go-bnb/initializers"
	"github.com/Arshad-Siddiqui/go-bnb/models"
)

func refreshUsers() {
	app = fiber.New()
	app.Get("/", controllers.UserIndex)
	app.Post("/", controllers.UserCreate)
	initializers.DB.Exec("DELETE FROM users")
}

var _ = Describe("UserCreate", func() {
	BeforeEach(refreshUsers)

	It("returns a 500 without valid params", func() {
		req, err := http.NewRequest("POST", "/", nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		Expect(resp.StatusCode).To(Equal(500))
	})

	It("returns a 201 Created with valid parameters", func() {
		req, err := http.NewRequest("POST", "/", nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Set("Content-Type", "application/json")

		user := models.User{
			Email:     "test@email.com",
			Password:  "test password",
			FirstName: "test first name",
			LastName:  "test last name",
		}

		json, _ := json.Marshal(user)

		req.Body = io.NopCloser(bytes.NewBuffer(json))
		resp, err := app.Test(req)
		if err != nil {
			fmt.Println(err)
			return
		}

		Expect(resp.StatusCode).To(Equal(http.StatusCreated))
	})

	It("Returns an error when an invalid email is provided", func() {
		req, err := http.NewRequest("POST", "/", nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Set("Content-Type", "application/json")

		user := models.User{
			Email:     "testemail.com",
			Password:  "test password",
			FirstName: "test first name",
			LastName:  "test last name",
		}

		json, _ := json.Marshal(user)

		req.Body = io.NopCloser(bytes.NewBuffer(json))
		resp, err := app.Test(req)
		if err != nil {
			fmt.Println(err)
			return
		}

		Expect(resp.StatusCode).To(Equal(550))
	})
})

var _ = Describe("UserIndex", func() {
	BeforeEach(refreshUsers)

	It("returns a 200 OK", func() {
		req, _ := http.NewRequest("GET", "/", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		Expect(resp.StatusCode).To(Equal(http.StatusOK))

	})

	It("returns a JSON response", func() {
		req, _ := http.NewRequest("GET", "/", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		Expect(resp.Header.Get("Content-Type")).To(Equal("application/json"))
	})
})
