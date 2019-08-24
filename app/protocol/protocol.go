package protocol

type Resp struct {
	Ret  int32       `json:"ret"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type AdminJwtTokenInfo struct {
	AdminId    int32
	Username   string
	Expiretime int32
}
