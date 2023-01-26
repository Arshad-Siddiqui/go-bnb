package controllers_test

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/Arshad-Siddiqui/go-bnb/controllers"
)

var _ = Describe("ListingIndex", func() {
	var app *fiber.App
	BeforeEach(func() {
		app = fiber.New()
		app.Get("/", controllers.ListingIndex)
	})

	It("returns a 200 OK", func() {
		req, _ := http.NewRequest("GET", "/", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		Expect(resp.StatusCode).To(Equal(http.StatusOK))
	})
})
