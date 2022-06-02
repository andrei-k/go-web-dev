package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/main", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	api := router.Group("/api")

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	router.Run()
}
