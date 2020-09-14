package login
//登录参数
type LoginRequest struct {
	Username string `json:"username" validate:"gt=4"`
	Password string `json:"password" validate:"gt=6"`
	Code     string `json:"code" validate:"len=6"`
}