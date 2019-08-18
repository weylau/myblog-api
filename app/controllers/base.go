package controllers

type Resp struct {
	Ret  int32       `json:"ret"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
