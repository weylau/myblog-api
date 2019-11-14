package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/weylau/myblog-api/app/protocol"
	"net/http"
)

type User struct {
}

//用户信息
type Info struct {
	Roles        []string `json:"roles"`
	Introduction string   `json:"introduction"`
	Avatar       string   `json:"avatar"`
	Name         string   `json:"name"`
}

//用户信息
func (User) Info(c *gin.Context) {
	resp := &protocol.Resp{Ret: 0, Msg: "", Data: ""}
	roles := []string{"admin"}
	resp.Data = &Info{
		Roles:  roles,
		Avatar: "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
		Name:   "admin",
	}
	c.JSON(http.StatusOK, resp)
}
