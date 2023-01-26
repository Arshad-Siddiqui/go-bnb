package controllers_test

import (
	"testing"

	"github.com/Arshad-Siddiqui/go-bnb/initializers"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestControllers(t *testing.T) {
	initializers.LoadTestEnvVariables()
	initializers.InitDB()

	RegisterFailHandler(Fail)
	RunSpecs(t, "Controllers Suite")
}
