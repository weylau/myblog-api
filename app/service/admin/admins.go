package admin

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"myblog-api/app/config"
	"myblog-api/app/db/mysql"
	"myblog-api/app/helper"
	"myblog-api/app/loger"
	"myblog-api/app/model"
	"myblog-api/app/protocol"
	"time"
	"github.com/juju/errors"
)

type Admins struct {
}


//用户信息
type UserInfo struct {
	AdminId  int32  `json:"admin_id"`
	UserName string `json:"user_name"`
	Token    string `json:"token"`
}

//登录
func (this *Admins) Login(username string, password string, code uint32) (resp protocol.Resp) {
	resp = protocol.Resp{Ret: -1, Msg: "", Data: ""}

	//校验谷歌验证码
	ga_code, err := helper.MkGaCode(config.Configs.GaSecret)
	if err != nil {
		loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
		resp.Msg = "系统错误"
		return resp
	}

	if code != ga_code {
		resp.Msg = "谷歌验证码错误"
		return resp
	}

	db := mysql.Default().GetConn()
	defer db.Close()
	//查询用户
	admin := model.Admins{}
	db.Where("username=?", username).First(&admin)
	if helper.IsEmpty(admin) {
		resp.Msg = "账号不存在"
		return resp
	}

	//检测密码是否正确
	if helper.MkMd5(password) != admin.Password {
		loger.Loger.Info("admin:", admin)
		loger.Loger.Info("服务端密码:", admin.Password)
		loger.Loger.Info("客户端密码:", helper.MkMd5(password))
		resp.Msg = "密码错误"
		return resp
	}

	//生成token
	token, err := helper.JwtEncode(jwt.MapClaims{"admin_id": fmt.Sprintf("%d", admin.AdminId), "username": admin.Username, "expr_time": fmt.Sprintf("%d", time.Now().Unix())}, []byte(config.Configs.JwtSecret))
	if err != nil {
		loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
		resp.Ret = -999
		resp.Msg = "系统错误"
		return resp
	}
	user_info := UserInfo{
		AdminId:  admin.AdminId,
		UserName: admin.Username,
		Token:    token,
	}
	resp.Data = user_info
	resp.Ret = 0
	return resp
}
