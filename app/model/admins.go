package model

type Admins struct {
	AdminId  int32  `json:"admin_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Status   int8   `json:"status"`
}

func (Admins) TableName() string {
	return "mb_admins"
}
