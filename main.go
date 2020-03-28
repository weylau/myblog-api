package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myblog-api/app"
	"myblog-api/app/config"
	"myblog-api/app/helper"
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
	application := app.Default()
	application.Run()
	err := application.GetEngin().Run("0.0.0.0:" + config.Configs.HttpListenPort)
	if err != nil {
		fmt.Println("start service error!!")
		os.Exit(0)
	}
}
