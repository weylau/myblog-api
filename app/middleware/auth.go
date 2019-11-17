package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myblog-api/app/config"
	"myblog-api/app/helper"
	"myblog-api/app/protocol"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Auth struct {
}

func Default() *Auth {
	return &Auth{}
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
		auth_info, err := helper.JwtDncode(token, []byte(config.Configs.JwtSecret))
		if err != nil {
			resp.Ret = 603
			resp.Msg = "登录失败，请重新登录" + err.Error()
			c.JSON(http.StatusOK, resp)
			c.Abort()
			return
		}
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

		if curent_time-auth_time > config.Configs.JwtExprTime {
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

func (Auth) Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			//c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma,x-token")
			//              允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                  //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")                                                                                                                                                              // 设置返回格式是json
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.Status(204)
			c.Abort()
			return
		}
		// 处理请求
		c.Next()
	}
}
