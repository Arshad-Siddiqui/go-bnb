package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	// c.String(http.StatusOK, "Hello World!")
	c.HTML(http.StatusOK, "index.html", nil)
}

func signup(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

func main() {
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})
	r.Static("/static", "./static/")
	r.LoadHTMLGlob("static/*.html")

	r.GET("/", index)
	r.GET("/signup", signup)
	r.Run(":8080")
}
