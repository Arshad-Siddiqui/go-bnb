package controllers_test

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/Arshad-Siddiqui/go-bnb/controllers"
	"github.com/Arshad-Siddiqui/go-bnb/initializers"
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
})
