package front

import (
	"github.com/weylau/myblog-api/app/db"
	"github.com/weylau/myblog-api/app/model"
)

type ArticleDetails struct {
	model.Articles
	Contents string `json:"contents"`
	ShowType int    `json:"show_type"`
}

type Articles struct {
}

/**
 *分页获取文章列表
 */
func (Articles) GetList(page int, page_size int, cate_id int, fields []string) ([]model.Articles, error) {
	db := db.DBConn()
	defer db.Close()
	offset := (page - 1) * page_size
	articles := make([]model.Articles, 0)
	if cate_id > 0 {
		db = db.Where("cate_id = ?", cate_id)
	}
	db.Select(fields).Offset(offset).Limit(page_size).Order("article_id desc").Find(&articles)
	return articles, nil
}

/**
获取文章详情
*/
func (Articles) GetArticleDetail(article_id int) *ArticleDetails {
	article_content := model.ArticlesContents{}
	article_details := ArticleDetails{}
	db := db.DBConn()
	defer db.Close()
	db.Where("article_id = ?", article_id).First(&article_details)
	db.Where("article_id = ?", article_id).First(&article_content)
	article_details.Contents = article_content.Contents
	article_details.ShowType = article_content.ShowType
	return &article_details
}

/**
获取文章类型
*/
func (Articles) GetArticleCate() []model.ArticlesCate {
	article_cates := make([]model.ArticlesCate, 0)
	db := db.DBConn()
	defer db.Close()
	db.Order("orderby asc").Find(&article_cates)
	return article_cates
}
