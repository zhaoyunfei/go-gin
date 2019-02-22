package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func main() {
	router := gin.Default()
	router.Any("/testing", startPage)
	router.Run(":8000")
}

func startPage(c *gin.Context)  {
	var person Person
	if c.ShouldBindQuery(&person) ==nil{
		log.Println("== only bind by query string==")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	c.String(200, "Success")
}
