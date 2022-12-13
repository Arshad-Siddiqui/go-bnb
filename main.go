package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func sayHello(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

func main() {
	r := gin.Default()
	r.GET("/", sayHello)
	r.Run(":8080")
}
