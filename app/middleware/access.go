package middleware

import (
	"github.com/gin-gonic/gin"
	"myblog-api/app/service/front"
	"time"
	"myblog-api/app/loger"
	"github.com/juju/errors"
)

func AddAccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		go func() {
			var accessLogServ front.AccessLog
			params := &front.AccessLogParams{}
			params.Ip = c.ClientIP()
			params.Timestamp = time.Now().Unix()
			params.Path = c.Request.Method + "|" + c.Request.URL.Path
			params.Date = time.Now().Format("2006-01-02")
			err := accessLogServ.Add(params)
			if err != nil {
				loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
			}
		}()
		c.Next()
	}

}
