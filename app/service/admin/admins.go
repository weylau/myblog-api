package admin

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/weylau/myblog-api/app/configs"
	"github.com/weylau/myblog-api/app/db"
	"github.com/weylau/myblog-api/app/helpers"
	"github.com/weylau/myblog-api/app/model"
	"github.com/weylau/myblog-api/app/protocol"
	"time"
)

type Admins struct {
}

type UserInfo struct {
	AdminId  int32  `json:"admin_id"`
	UserName string `json:"user_name"`
	Token    string `json:"token"`
}

func (Admins) Login(username string, password string) (resp protocol.Resp) {
	resp = protocol.Resp{Ret: -1, Msg: "", Data: ""}
	helper := helpers.Helpers{}

	db := db.DBConn()
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
		fmt.Println("admin:", admin)
		fmt.Println("服务端：", admin.Password)
		fmt.Println("客户端：", helper.MkMd5(password))
		resp.Msg = "密码错误"
		return resp
	}

	//生成token
	token, err := helper.JwtEncode(jwt.MapClaims{"admin_id": admin.AdminId, "username": admin.Username, "expr_time": time.Now().Unix()}, []byte(configs.JwtSecret))
	if err != nil {
		resp.Ret = -999
		resp.Msg = "系统错误:" + err.Error()
		return resp
	}
	fmt.Println("token:", token)
	user_info := UserInfo{
		AdminId:  admin.AdminId,
		UserName: admin.Username,
		Token:    token,
	}
	resp.Data = user_info
	resp.Ret = 0
	return resp
}
