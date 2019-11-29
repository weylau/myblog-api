package helper

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base32"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"time"
)

func MkMd5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func IsEmpty(params interface{}) bool {
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
func JwtEncode(jwtinfo jwt.MapClaims, secret_key []byte) (jwt_token string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtinfo)
	tokenString, err := token.SignedString(secret_key)
	return tokenString, err
}

//jwt解密
func JwtDncode(token_string string, secret_key interface{}) (token_info map[string]interface{}, err error) {
	token, err := jwt.Parse(token_string, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secret_key, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

func Interface2String(inter interface{}) (str string) {
	str = ""
	switch inter.(type) {
	case string:
		str = inter.(string)
	default:

	}
	return str
}

/**
 * 获取谷歌验证码
 */
func MkGaCode(secret string) (code uint32, err error) {

	// decode the key from the first argument
	inputNoSpaces := strings.Replace(secret, " ", "", -1)
	inputNoSpacesUpper := strings.ToUpper(inputNoSpaces)
	fmt.Println(inputNoSpacesUpper)
	key, err := base32.StdEncoding.DecodeString(inputNoSpacesUpper)
	if err != nil {
		return 0, err
	}

	// generate a one-time password using the time at 30-second intervals
	epochSeconds := time.Now().Unix()
	code = oneTimePassword(key, toBytes(epochSeconds/30))

	return code, nil
}

func oneTimePassword(key []byte, value []byte) uint32 {
	// sign the value using HMAC-SHA1
	hmacSha1 := hmac.New(sha1.New, key)
	hmacSha1.Write(value)
	hash := hmacSha1.Sum(nil)

	// We're going to use a subset of the generated hash.
	// Using the last nibble (half-byte) to choose the index to start from.
	// This number is always appropriate as it's maximum decimal 15, the hash will
	// have the maximum index 19 (20 bytes of SHA1) and we need 4 bytes.
	offset := hash[len(hash)-1] & 0x0F

	// get a 32-bit (4-byte) chunk from the hash starting at offset
	hashParts := hash[offset : offset+4]

	// ignore the most significant bit as per RFC 4226
	hashParts[0] = hashParts[0] & 0x7F

	number := toUint32(hashParts)

	// size to 6 digits
	// one million is the first number with 7 digits so the remainder
	// of the division will always return < 7 digits
	pwd := number % 1000000

	return pwd
}

func toBytes(value int64) []byte {
	var result []byte
	mask := int64(0xFF)
	shifts := [8]uint16{56, 48, 40, 32, 24, 16, 8, 0}
	for _, shift := range shifts {
		result = append(result, byte((value>>shift)&mask))
	}
	return result
}

func toUint32(bytes []byte) uint32 {
	return (uint32(bytes[0]) << 24) + (uint32(bytes[1]) << 16) +
		(uint32(bytes[2]) << 8) + uint32(bytes[3])
}

func IsTimeStr(str string) bool {
	timeLayout := "2006-01-02 15:04:05"                        //转化所需模板
	loc, _ := time.LoadLocation("Local")                       //重要：获取时区
	theTime, err := time.ParseInLocation(timeLayout, str, loc) //使用模板在对应时区转化为time.time类型
	if err != nil {
		return false
	}
	if theTime.Unix() > 0 {
		return true
	}
	return false
}

//时间格式转换
func DateToDateTime(date string) string {
	timeTemplate := "2006-01-02T15:04:05+08:00" //常规类型
	toTemplate := "2006-01-02 15:04:05"
	stamp, _ := time.ParseInLocation(timeTemplate, date, time.Local)
	return time.Unix(stamp.Unix(), 0).Format(toTemplate)

}

func GetAppDir() string {
	appDir, err := os.Getwd()
	if err != nil {
		file, _ := exec.LookPath(os.Args[0])
		applicationPath, _ := filepath.Abs(file)
		appDir, _ = filepath.Split(applicationPath)
	}
	return appDir
}
