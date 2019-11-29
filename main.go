package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myblog-api/app/config"
	"myblog-api/app/helper"
	"myblog-api/app/router"
	"os"
)

func init() {
	appDir := helper.GetAppDir()
	config.Default(appDir + "/config.ini")

}

func main() {
	if config.Configs.Env == "prd" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := router.Default()
	r.Run()
	err := r.GetEngin().Run("0.0.0.0:" + config.Configs.HttpListenPort)
	if err != nil {
		fmt.Println("start service error!!")
		os.Exit(0)
	}
}
