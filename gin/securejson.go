package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/somejson", func(c *gin.Context) {
		names := []string{"lena", "austin", "foo"}

		c.SecureJSON(http.StatusOK, names)
	})
	r.Run(":8080")
}
