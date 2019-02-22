package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	//日志记录方式：
	//禁用控制台颜色
	gin.DisableConsoleColor()
	//创建记录日志的文件
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout) //将日志同时写入文件和控制台
	//默认带中间件启动方式，Use(Logger(), Recovery()
	//也可以自己New()，使用use添加中间件
	router := gin.Default()
	//添加自定义日志文件格式
	router.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			params.ClientIP,
			params.TimeStamp.Format(time.RFC1123),
			params.Method,
			params.Path,
			params.Request.Proto,
			params.StatusCode,
			params.Latency,
			params.Request.UserAgent(),
			params.ErrorMessage,)
	}))
	//router.GET("/someGet", getting)
	//router.POST("/somePost", posting)
	//router.PUT("somePut", putting)
	//router.DELETE("someDelete", deleting)
	//router.PATCH("somePatch", patching)
	//router.HEAD("someHead", head)
	//router.OPTIONS("someOptions", options)
	//router.Run()
	//路由分组
	v1 := router.Group("/v1")
	{
		//匹配url：/user/zhaohe
		v1.GET("/user/:name", func(c *gin.Context) {
			name := c.Param("name")
			c.String(http.StatusOK, "Hello %s", name)
		})
		//匹配url： /user/zhaohe/send
		v1.GET("/user/:name/*action", func(c *gin.Context) {
			name := c.Param("name")
			action := c.Param("action")
			message := name + " is " + action
			c.String(http.StatusOK, message)
		})
	}
	//获取get参数，匹配url: /welcome?firstname=he&lastname=zhao
	router.GET("/welcome", func(c *gin.Context) {
		firstName := c.DefaultQuery("firstname", "Guest")
		lastName := c.Query("lastname")

		c.String(http.StatusOK, "hello %s %s", firstName, lastName)
	})
	//获取post参数
	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonnymous")

		c.JSON(200, gin.H{
			"status": "posted",
			"message": message,
			"nick": nick,
		})
	})
	//get+post请求
	router.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s, message: %s", id, page, name, message)
	})
	//上传文件，测试：curl -X POST http://localhost:8000/upload -F "file=@Users/XXX/XXX.txt" -H "Content-Type: multipart/form-data"
	router.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		//上传文件到指定目录
		//c.SaveUploadedFile(file, "dst")
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploadede!", file.Filename))
	})
	//多文件上传，测试: curl -X POST http://localhost:8000/multiupload -F "upload[]=@README.md" -F "upload[]=@requirements.txt" -H "Content-Type: multipart/form-data"
	router.POST("/multiupload", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]
		for _, file := range files{
			log.Println(file.Filename)
			//c.SaveUploadedFile(file, "dst")
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploads!", len(files)))
	})
	router.Run(":8000")
}

