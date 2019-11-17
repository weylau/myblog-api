package admin

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"myblog-api/app/protocol"
	"myblog-api/app/service/admin"
	"io/ioutil"
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

//登录
func (Login) Login(c *gin.Context) {
	resp := protocol.Resp{Ret: 0, Msg: "", Data: ""}
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		resp.Ret = -1
		resp.Msg = "用户名或密码不能为空1"
		c.JSON(http.StatusOK, resp)
		return
	}
	loginParams := &LoginParams{}
	err = json.Unmarshal(data, loginParams)
	if err != nil {
		resp.Ret = -1
		resp.Msg = "用户名或密码不能为空2" + err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	username := loginParams.Username
	password := loginParams.Password
	code, err := strconv.Atoi(loginParams.Code)
	if err != nil {
		resp.Ret = -1
		resp.Msg = "谷歌验证码错误：" + err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	admin_serv := admin.Admins{}
	validator,_ := validate.Default()
	if check := validator.CheckStruct(loginParams); !check{
		resp.Ret = -1
		resp.Msg = validator.GetOneError()
		c.JSON(http.StatusOK, resp)
		return
	}

	resp = admin_serv.Login(username, password, uint32(code))
	c.JSON(http.StatusOK, resp)
}
