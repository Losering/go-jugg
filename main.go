package main

import (
	"jugg/jugg"
	"net/http"
)

func main() {
	r := jugg.New()

	r.GET("/", func(ctx *jugg.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello jugg</h1>")
	})
	r.GET("/hello", func(c *jugg.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	r.GET("/hello/:name", func(c *jugg.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})
	r.GET("/hello/:name/hello", func(c *jugg.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *jugg.Context) {
		c.JSON(http.StatusOK, jugg.H{"filepath": c.Param("filepath")})
	})

	r.POST("/login", func(c *jugg.Context) {
		c.JSON(http.StatusOK, jugg.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	r.Run(":9999")
}
