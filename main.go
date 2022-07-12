package main

import (
	"jugg/jugg"
	"net/http"
)

func main() {
	r := jugg.New()

	r.GET("/", func(ctx *jugg.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(c *jugg.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	r.POST("/login", func(c *jugg.Context) {
		c.JSON(http.StatusOK, jugg.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	r.Run(":9999")
}
