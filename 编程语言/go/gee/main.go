package main

import (
	"gee/gee"
	"net/http"
)

func main() {
	r := gee.Default()
	r.GET("/", func(c *gee.Context) {

		c.String(http.StatusOK, "Hello Geektutu\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *gee.Context) {

		names := []string{"geektutu"}
		c.String(http.StatusOK, names[1])

	})

	r.GET("/test", func(c *gee.Context) {
		c.JSON(http.StatusOK, "aa")
	})

	r.Run(":9999")
}
