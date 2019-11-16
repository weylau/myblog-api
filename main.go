package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/weylau/myblog-api/app/config"
	"github.com/weylau/myblog-api/app/router"
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
	config.SetUp(appDir + "/config.ini")
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
