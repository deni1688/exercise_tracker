package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/exercises", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "notImplemented",
		})
	})

	r.POST("/exercises", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "notImplemented",
		})
	})

	r.Run()
}
