package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
	"net/http"
	"reflect"
	"time"
)

//自定义验证器
type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

func bookableDate(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,) bool {
	if date, ok := field.Interface().(time.Time); ok{
		today := time.Now()
		if today.Year() > date.Year() || today.YearDay() > date.YearDay() {
			return false
		}
	}
	return true
}

func main() {
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookableDate)
	}
	router.GET("/bookable", getBookable)
	router.Run(":8000")
}

func getBookable(c *gin.Context)  {
	var b Booking
	if err := c.ShouldBindWith(&b, binding.Query); err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Booking dates ate valid"})
	}else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
