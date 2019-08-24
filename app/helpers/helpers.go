package helpers

import (
	"crypto/md5"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"reflect"
)

type Helpers struct {
}

func (Helpers) MkMd5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func (Helpers) IsEmpty(params interface{}) bool {
	//初始化变量
	var (
		flag          bool = true
		default_value reflect.Value
	)

	r := reflect.ValueOf(params)

	//获取对应类型默认值
	default_value = reflect.Zero(r.Type())
	//由于params 接口类型 所以default_value也要获取对应接口类型的值 如果获取不为接口类型 一直为返回false
	if !reflect.DeepEqual(r.Interface(), default_value.Interface()) {
		flag = false
	}
	return flag
}

//jwt加密
func (Helpers) JwtEncode(jwtinfo jwt.MapClaims, secret_key []byte) (jwt_token string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtinfo)
	tokenString, err := token.SignedString(secret_key)
	return tokenString, err
}

//jwt加密
func (Helpers) JwtDncode(token_string string, secret_key interface{}) {
	token, err := jwt.Parse(token_string, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return secret_key, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims)
	} else {
		fmt.Println(err)
	}
}
