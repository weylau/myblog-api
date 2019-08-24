package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/weylau/myblog-api/app/protocol"
	"github.com/weylau/myblog-api/app/service/admin"
	"net/http"
)

type Login struct {
}

func (Login) Login(c *gin.Context) {
	resp := protocol.Resp{Ret: 0, Msg: "", Data: ""}
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")
	admin_serv := admin.Admins{}
	if username == "" || password == "" {
		resp.Ret = -1
		resp.Msg = "用户名或密码不能为空"
		c.JSON(http.StatusOK, resp)
		return
	}

	resp = admin_serv.Login(username, password)
	c.JSON(http.StatusOK, resp)
}
