package front

import (
	"github.com/weylau/myblog-api/app/db"
	"github.com/weylau/myblog-api/app/model"
)

type ArticleDetails struct {
	model.Articles
	Contents string `json:"contents"`
}

type Articles struct {
}

/**
 *分页获取文章列表
 */
func (Articles) GetList(page int, page_size int, fields []string) ([]model.Articles, error) {
	db := db.DBConn()
	defer db.Close()
	offset := (page - 1) * page_size
	articles := make([]model.Articles, 0)
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
	return &article_details
}
