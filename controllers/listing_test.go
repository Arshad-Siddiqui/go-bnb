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

var app *fiber.App

func refreshListings() {
	app = fiber.New()
	app.Get("/", controllers.ListingIndex)
	app.Post("/", controllers.ListingCreate)
	initializers.DB.Exec("DELETE FROM listings")
}

var _ = Describe("ListingIndex", func() {
	BeforeEach(func() {
		refreshListings()
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

var _ = Describe("ListingCreate", func() {
	BeforeEach(func() {
		refreshListings()
	})

	It("returns a 201 Created", func() {
		req, err := http.NewRequest("POST", "/", nil)
		if err != nil {
			fmt.Println(err)
			return
		}

		req.Header.Set("Content-Type", "application/json")

		listing := models.Listing{
			Title:       "Test Listing",
			Description: "Test Description",
			Price:       100,
		}

		jsonValue, err := json.Marshal(listing)
		if err != nil {
			fmt.Println(err)
			return
		}

		req.Body = io.NopCloser(bytes.NewBuffer(jsonValue))
		resp, err := app.Test(req)
		if err != nil {
			fmt.Println(err)
			return
		}

		Expect(resp.StatusCode).To(Equal(http.StatusCreated))
	})

	It("returns a JSON response", func() {
		req, err := http.NewRequest("POST", "/", nil)
		if err != nil {
			fmt.Println(err)
			return
		}

		req.Header.Set("Content-Type", "application/json")

		listing := models.Listing{
			Title:       "Test Listing",
			Description: "Test Description",
			Price:       100,
		}

		jsonValue, err := json.Marshal(listing)
		if err != nil {
			fmt.Println(err)
			return
		}

		req.Body = io.NopCloser(bytes.NewBuffer(jsonValue))
		resp, err := app.Test(req)
		if err != nil {
			fmt.Println(err)
			return
		}

		Expect(resp.Header.Get("Content-Type")).To(Equal("application/json"))
	})

	It("Should return the listing created", func() {
		req, err := http.NewRequest("POST", "/", nil)
		if err != nil {
			fmt.Println(err)
			return
		}

		req.Header.Set("Content-Type", "application/json")

		listing := models.Listing{
			Title:       "Test Listing",
			Description: "Test Description",
			Price:       100,
		}

		jsonValue, err := json.Marshal(listing)
		if err != nil {
			fmt.Println(err)
			return
		}

		req.Body = io.NopCloser(bytes.NewBuffer(jsonValue))
		resp, err := app.Test(req)
		if err != nil {
			fmt.Println(err)
			return
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		Expect(string(body)).To(Equal(`{"id":1,"title":"Test Listing","description":"Test Description","price":100}`))
	})
})
