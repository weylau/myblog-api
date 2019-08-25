package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/weylau/myblog-api/app/configs"
	"github.com/weylau/myblog-api/app/helpers"
	"github.com/weylau/myblog-api/app/protocol"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Auth struct {
}

func (Auth) CheckAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		resp := protocol.Resp{Ret: -1, Msg: "", Data: ""}
		authorization := c.Request.Header.Get("Authorization")
		if authorization == "" {
			resp.Ret = 601
			resp.Msg = "登录失败，请重新登录"
			c.JSON(http.StatusOK, resp)
			c.Abort()
			return
		}
		fmt.Println("authorization:", authorization)
		kv := strings.Split(authorization, " ")
		if len(kv) != 2 || kv[0] != "Bearer" {
			resp.Ret = 602
			resp.Msg = "登录失败，请重新登录"
			c.JSON(http.StatusOK, resp)
			c.Abort()
			return
		}
		token := kv[1]
		fmt.Println("token:", token)
		helper := helpers.Helpers{}
		auth_info, err := helper.JwtDncode(token, []byte(configs.JwtSecret))
		if err != nil {
			resp.Ret = 603
			resp.Msg = "登录失败，请重新登录" + err.Error()
			c.JSON(http.StatusOK, resp)
			c.Abort()
			return
		}
		fmt.Println("auth_info:", auth_info)
		curent_time := time.Now().Unix()
		admin_id := helper.Interface2String(auth_info["admin_id"])
		username := helper.Interface2String(auth_info["username"])
		expr_time := helper.Interface2String(auth_info["expr_time"])
		if admin_id == "" || username == "" || expr_time == "" {
			resp.Ret = 605
			resp.Msg = "登录失败，请重新登录"
			c.JSON(http.StatusOK, resp)
			c.Abort()
			return
		}
		auth_time, err := strconv.ParseInt(expr_time, 10, 64)
		if err != nil {
			resp.Ret = 606
			resp.Msg = "登录失败，请重新登录" + err.Error()
			c.JSON(http.StatusOK, resp)
			c.Abort()
			return
		}

		if curent_time-auth_time > configs.JwtExprTime {
			resp.Ret = 607
			resp.Msg = "登录失效，请重新登录"
			c.JSON(http.StatusOK, resp)
			c.Abort()
			return
		}
		c.Set("admin_id", admin_id)
		c.Set("username", username)
		c.Next()
	}

}
