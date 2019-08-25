package main

import (
	"github.com/gin-gonic/gin"
	"github.com/weylau/myblog-api/app/controllers/admin"
	"github.com/weylau/myblog-api/app/controllers/front"
	"github.com/weylau/myblog-api/app/middleware"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	auth_middleware := middleware.Auth{}
	r.Use(auth_middleware.Cors())

	article_front_ctrl := front.Articles{}
	//文章相关
	r.GET("/article/list", article_front_ctrl.GetList)
	r.GET("/article/detail", article_front_ctrl.GetDetail)

	//后台管理
	login_admin_ctrl := admin.Login{}
	article_admin_ctrl := admin.Articles{}
	r.POST("/adapi/login", login_admin_ctrl.Login)
	authorized := r.Group("/adapi")
	authorized.Use(auth_middleware.CheckAuth())
	{
		authorized.POST("/articles/add", article_admin_ctrl.Add)
	}

	return r
}

func main() {
	r := setupRouter()
	r.Run("0.0.0.0:8080")
}
