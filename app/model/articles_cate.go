package model

type ArticlesCate struct {
	CateId int    `json:"cate_id"`
	Name   string `json:"name"`
	CName  string `json:"c_name"`
}

func (ArticlesCate) TableName() string {
	return "mb_articles_cate"
}
