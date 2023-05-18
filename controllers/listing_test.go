package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/Arshad-Siddiqui/go-bnb/controllers"
	"github.com/Arshad-Siddiqui/go-bnb/initializers"
	"github.com/Arshad-Siddiqui/go-bnb/models"
	"github.com/Arshad-Siddiqui/jsonRequest"
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
		listing := models.Listing{
			Title:       "Test Listing",
			Description: "Test Description",
			Price:       100,
		}
		req, err := jsonRequest.PostJsonReq("/", listing)
		if err != nil {
			fmt.Println(err)
			return
		}

		resp, err := app.Test(req)
		if err != nil {
			fmt.Println(err)
			return
		}

		Expect(resp.StatusCode).To(Equal(http.StatusCreated))
	})

	It("returns a JSON response", func() {
		listing := models.Listing{
			Title:       "Test Listing",
			Description: "Test Description",
			Price:       100,
		}

		req, err := jsonRequest.PostJsonReq("/", listing)
		if err != nil {
			fmt.Println(err)
			return
		}

		resp, err := app.Test(req)
		if err != nil {
			fmt.Println(err)
			return
		}

		Expect(resp.Header.Get("Content-Type")).To(Equal("application/json"))
	})

	It("Should return the listing created", func() {
		listing := models.Listing{
			Title:       "Test Listing",
			Description: "Test Description",
			Price:       100,
		}

		req, err := jsonRequest.PostJsonReq("/", listing)
		if err != nil {
			fmt.Println(err)
			return
		}

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

	It("Should work multiple times", func() {
		req, err := http.NewRequest("POST", "/", nil)
		if err != nil {
			fmt.Println(err)
			return
		}

		req.Header.Set("Content-Type", "application/json")

		for i := 0; i < 3; i++ {
			listing := models.Listing{
				Title:       "Test " + strconv.Itoa(i),
				Description: "Test Description " + strconv.Itoa(i),
				Price:       100,
			}

			jsonValue, err := json.Marshal(listing)
			if err != nil {
				fmt.Println(err)
				return
			}

			req.Body = io.NopCloser(bytes.NewBuffer(jsonValue))
			_, err = app.Test(req)
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		req, _ = http.NewRequest("GET", "/", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		body, _ := io.ReadAll(resp.Body)

		listings := []models.Listing{}
		json.Unmarshal(body, &listings)
		Expect(len(listings)).To(Equal(3))
		Expect(listings[0].Title).To(Equal("Test 0"))
		Expect(listings[1].Description).To(Equal("Test Description 1"))
		Expect(listings[2].Price).To(Equal(100))
	})
})

var _ = Describe("ListingDelete", func() {
	BeforeEach(func() {
		refreshListings()
	})
})
