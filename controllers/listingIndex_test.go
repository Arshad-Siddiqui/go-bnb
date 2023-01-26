package controllers_test

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/Arshad-Siddiqui/go-bnb/controllers"
	"github.com/Arshad-Siddiqui/go-bnb/initializers"
)

var _ = Describe("ListingIndex", func() {
	var app *fiber.App
	BeforeEach(func() {
		app = fiber.New()
		app.Get("/", controllers.ListingIndex)
		app.Post("/", controllers.ListingCreate)
		initializers.DB.Exec("DELETE FROM listings")
	})

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
