package model

import (
	"github.com/weylau/myblog-api/app/db"
)

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

/**
 *分页获取文章列表
 */
func (Articles) GetList(page int, page_size int, fields []string) ([]Articles, error) {
	db := db.DBConn()
	defer db.Close()
	offset := (page - 1) * page_size
	articles := make([]Articles, 0)
	db.Select(fields).Offset(offset).Limit(page_size).Order("article_id desc").Find(&articles)
	return articles, nil
}
