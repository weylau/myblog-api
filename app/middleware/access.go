package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"myblog-api/app/loger"
	"myblog-api/app/service/front"
	"time"
)

func AddAccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params front.AccessLogParams
		var accessLogServ front.AccessLog
		params.Ip = c.ClientIP()
		params.Timestamp = time.Now().Unix()
		params.Path = c.Request.Method + "|" + c.Request.URL.Path
		err := accessLogServ.Add(params)
		paramsstr, _ := json.Marshal(params)
		loger.Default().Info("AddAccessLog-params:", string(paramsstr))
		if err != nil {
			loger.Default().Error("AddAccessLog-error:", err.Error())
		}
		c.Next()
	}

}
