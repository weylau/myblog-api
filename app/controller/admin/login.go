package admin

import (
	"github.com/gin-gonic/gin"
	"myblog-api/app/loger"
	"myblog-api/app/protocol"
	"myblog-api/app/request/admin/login"
	"myblog-api/app/service/admin"
	"myblog-api/app/validate"
	"net/http"
	"strconv"
	"github.com/juju/errors"
)

type Login struct {
}

//登录
func (this *Login) Login(c *gin.Context) {
	resp := protocol.Resp{Ret: 0, Msg: "", Data: ""}
	var loginRequest login.LoginRequest
	err := c.ShouldBindJSON(&loginRequest)
	if err != nil {
		loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
		resp.Ret = -1
		resp.Msg = "用户名或密码不能为空1"
		c.JSON(http.StatusOK, resp)
		return
	}
	username := loginRequest.Username
	password := loginRequest.Password
	code, err := strconv.Atoi(loginRequest.Code)
	if err != nil {
		loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
		resp.Ret = -1
		resp.Msg = "谷歌验证码错误"
		c.JSON(http.StatusOK, resp)
		return
	}

	serv := admin.Admins{}
	validator, _ := validate.Default()
	if check := validator.CheckStruct(loginRequest); !check {
		resp.Ret = -1
		resp.Msg = validator.GetOneError()
		c.JSON(http.StatusOK, resp)
		return
	}

	resp = serv.Login(username, password, uint32(code))
	c.JSON(http.StatusOK, resp)
}
