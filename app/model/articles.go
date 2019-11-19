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
	Status      int    `json:"status"`
}

func (Articles) TableName() string {
	return "mb_articles"
}
