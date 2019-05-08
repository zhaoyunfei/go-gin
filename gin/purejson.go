package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>hello, world.</b>",
		})
	})

	r.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>hello, world.</b>",
		})
	})
	r.Run(":8080")
}
