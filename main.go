package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/weylau/myblog-api/app/configs"
	"github.com/weylau/myblog-api/app/controllers/admin"
	"github.com/weylau/myblog-api/app/controllers/front"
	"github.com/weylau/myblog-api/app/middleware"
	"github.com/weylau/myblog-api/app/protocol"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func init() {
	var err error
	appDir, err := os.Getwd()
	if err != nil {
		file, _ := exec.LookPath(os.Args[0])
		applicationPath, _ := filepath.Abs(file)
		appDir, _ = filepath.Split(applicationPath)
	}
	configs.SetUp(appDir + "/config.ini")
}

func setupRouter() *gin.Engine {
	if configs.Configs.Env == "prd" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	auth_middleware := middleware.Auth{}
	r.Use(auth_middleware.Cors())

	article_front_ctrl := front.Articles{}
	//文章相关
	r.GET("/article/list", article_front_ctrl.GetList)
	r.GET("/article/detail", article_front_ctrl.GetDetail)
	r.GET("/article/cate", article_front_ctrl.GetCate)

	//后台管理
	login_admin_ctrl := admin.Login{}
	article_admin_ctrl := admin.Articles{}
	user_admin_ctrl := admin.User{}
	r.POST("/adapi/login", login_admin_ctrl.Login)
	authorized := r.Group("/adapi")
	authorized.Use(auth_middleware.CheckAuth())
	{
		authorized.POST("/article/add", article_admin_ctrl.Add)
		authorized.GET("/article/list", article_admin_ctrl.GetList)
		authorized.DELETE("/article/:id", article_admin_ctrl.Delete)
		authorized.GET("/article/detail/:id", article_admin_ctrl.Detail)
		authorized.GET("/user/info", user_admin_ctrl.Info)
	}
	//404
	r.NoRoute(resp404)

	return r
}

//404
func resp404(c *gin.Context) {
	resp := protocol.Resp{Ret: 404, Msg: "page not exists!", Data: ""}
	//返回404状态码
	c.JSON(http.StatusNotFound, resp)
}

func main() {
	r := setupRouter()

	err := r.Run("0.0.0.0:" + configs.Configs.HttpListenPort)
	if err != nil {
		fmt.Println("http服务启动失败")
		os.Exit(0)
	}
}
