package main

import (
	"github.com/gin-gonic/gin"
	"github.com/weylau/myblog-api/app/controllers/front"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	//文章相关
	r.GET("/article/list", front.GetList)
	r.GET("/article/detail", front.GetDetail)

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
