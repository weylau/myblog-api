package model

type Articles struct {
	ArticleId   int    `json:"article_id",gorm:"primary_key;AUTO_INCREMENT"`
	CateId      int    `json:"cate_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Keywords    string `json:"keywords"`
	ImgPath     string `json:"img_path"`
	ModifyTime  string `json:"modify_time"`
	OpId        int    `json:"op_id"`
	OpUser      string `json:"op_user"`
	CreateTime  string `json:"create_time"`
}

func (Articles) TableName() string {
	return "mb_articles"
}

func (this Articles) GetCateName() string {
	cate_name := ""
	switch this.CateId {
	case 1:
		cate_name = "php"
	case 2:
		cate_name = "golang"
	case 3:
		cate_name = "mysql"
	default:
		cate_name = ""
	}
	return cate_name
}
