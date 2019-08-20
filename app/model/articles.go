package model

type Articles struct {
	ArticleId   int32  `json:"article_id"`
	CateId      int32  `json:"cate_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Keywords    string `json:"keywords"`
	ImgPath     string `json:"img_path"`
	ModifyTime  string `json:"modify_time"`
	OpId        int32  `json:"op_id"`
	OpUser      string `json:"op_user"`
	CreateTime  string `json:"create_time"`
}

func (Articles) TableName() string {
	return "mb_articles"
}
