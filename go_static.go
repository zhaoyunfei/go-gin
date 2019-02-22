package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.Static("/assets", "./assets") //使用本地路径
	router.StaticFS("/more_static", http.Dir("file path")) // 网络地址
	router.StaticFile("/favicon.ico", "./resources/favicon.ico") //文件地址
}
