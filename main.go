package main

import (
	"ges"
	"net/http"
)

func main() {
	r := ges.New()
	r.GET("/", func(c *ges.Context) {
		c.HTML(http.StatusOK, "<h1>Hello ges</h1>")
	})
	r.GET("/hello", func(c *ges.Context) {
		// expect /hello?name=gesktutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *ges.Context) {
		c.JSON(http.StatusOK, ges.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.RUN(":9999")
}
