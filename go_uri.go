package main

import "github.com/gin-gonic/gin"

type PersonUri struct{
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func main() {
	router := gin.Default()
	router.GET("/:name/:id", func(c *gin.Context) {
		var personUri PersonUri
		if err := c.ShouldBindUri(&personUri); err != nil{
			c.JSON(400, gin.H{"msg": err})
			return
		}
		c.JSON(200, gin.H{"name": personUri.Name, "uuid": personUri.ID})
	})
	router.Run(":8000")
}
