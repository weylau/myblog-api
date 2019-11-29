package admin

import (
	"github.com/gin-gonic/gin"
	"myblog-api/app/loger"
	"myblog-api/app/protocol"
	"myblog-api/app/service/admin"
	"myblog-api/app/validate"
	"net/http"
	"strconv"
)

type Login struct {
}

//登录参数
type LoginParams struct {
	Username string `json:"username" validate:"gt=4"`
	Password string `json:"password" validate:"gt=6"`
	Code     string `json:"code" validate:"len=6"`
}

func (*Login) getLogTitle() string {
	return "ctrller-admin-login-"
}

//登录
func (this *Login) Login(c *gin.Context) {
	resp := protocol.Resp{Ret: 0, Msg: "", Data: ""}
	var loginParams LoginParams
	err := c.ShouldBindJSON(&loginParams)
	if err != nil {
		loger.Default().Error(this.getLogTitle(), "Login-error1:", err.Error())
		resp.Ret = -1
		resp.Msg = "用户名或密码不能为空1"
		c.JSON(http.StatusOK, resp)
		return
	}
	username := loginParams.Username
	password := loginParams.Password
	code, err := strconv.Atoi(loginParams.Code)
	if err != nil {
		loger.Default().Error(this.getLogTitle(), "Login-error2:", err.Error())
		resp.Ret = -1
		resp.Msg = "谷歌验证码错误"
		c.JSON(http.StatusOK, resp)
		return
	}

	admin_serv := admin.Admins{}
	validator, _ := validate.Default()
	if check := validator.CheckStruct(loginParams); !check {
		resp.Ret = -1
		resp.Msg = validator.GetOneError()
		c.JSON(http.StatusOK, resp)
		return
	}

	resp = admin_serv.Login(username, password, uint32(code))
	c.JSON(http.StatusOK, resp)
}
