package router

import (
	"github.com/gin-gonic/gin"
	"myblog-api/app/controller/admin"
	"myblog-api/app/controller/front"
	"myblog-api/app/middleware"
	"myblog-api/app/protocol"
	"net/http"
)

type Router struct {
	engine *gin.Engine
}

func Default() *Router {
	router := &Router{}
	router.engine = gin.Default()
	return router
}

func (this *Router) Run() {
	this.SetCors()
	this.setFront()
	this.setAdmin()
	this.set404()
}

func (this *Router) GetEngin() *gin.Engine {
	return this.engine
}

func (this *Router) SetCors() {
	this.engine.Use(middleware.Cors())
}

func (this *Router) setFront() {
	article_front_ctrl := front.Articles{}
	this.engine.GET("/articles", article_front_ctrl.GetList)
	this.engine.GET("/categories", article_front_ctrl.GetCategories)
	this.engine.GET("/article/:id", article_front_ctrl.Show)
}

func (this *Router) setAdmin() {
	//后台管理
	login_admin_ctrl := admin.Login{}
	article_admin_ctrl := admin.Articles{}
	user_admin_ctrl := admin.User{}
	this.engine.POST("/adapi/login", login_admin_ctrl.Login)
	authorized := this.engine.Group("/adapi")
	authorized.Use(middleware.CheckAuth())
	{
		authorized.POST("/article", article_admin_ctrl.Add)
		authorized.PUT("/article/:id", article_admin_ctrl.Update)
		authorized.GET("/articles", article_admin_ctrl.GetList)
		authorized.DELETE("/article/:id", article_admin_ctrl.Delete)
		authorized.GET("/article/:id", article_admin_ctrl.Show)
		authorized.GET("/user", user_admin_ctrl.Show)
	}
}

func (this *Router) set404() {
	this.engine.NoRoute(func(context *gin.Context) {
		resp := protocol.Resp{Ret: 404, Msg: "page not exists!", Data: ""}
		//返回404状态码
		context.JSON(http.StatusNotFound, resp)
	})
}
