package admin

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/weylau/myblog-api/app/protocol"
	"github.com/weylau/myblog-api/app/service/admin"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Login struct {
}

//登录参数
type LoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Code     string `json:"code"`
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
	if username == "" || password == "" {
		resp.Ret = -1
		resp.Msg = "用户名或密码不能为空"
		c.JSON(http.StatusOK, resp)
		return
	}

	resp = admin_serv.Login(username, password, uint32(code))
	c.JSON(http.StatusOK, resp)
}
