package front

import (
	"context"
	"fmt"
	"myblog-api/app/db/es"
	"myblog-api/app/db/mysql"
	"myblog-api/app/helper"
	"myblog-api/app/model"
	"github.com/olivere/elastic"
	"reflect"
	//"reflect"
)

//文章详情
type ArticleDetails struct {
	model.Articles
	Contents string `json:"contents"`
	ShowType int    `json:"show_type"`
}

type Articles struct {
}

//分页获取文章列表mysql
func (Articles) GetListForMysql(page int, page_size int, cate_id int, fields []string) ([]model.Articles, error) {
	db := mysql.Default().GetConn()
	defer db.Close()
	offset := (page - 1) * page_size
	articles := make([]model.Articles, 0)
	if cate_id > 0 {
		db = db.Where("cate_id = ?", cate_id)
	}
	db.Select(fields).Offset(offset).Limit(page_size).Order("article_id desc").Find(&articles)
	return articles, nil
}

//分页获取文章列表es
func (Articles) GetListForEs(page int, page_size int, cate_id int, fields []string) ([]model.Articles, error) {
	esconn := es.Default().GetConn()
	ctx := context.Background()
	articles := make([]model.Articles, 0)
	query := esconn.Search().
		Index("myblog").
		Type("mb_articles").
		Size(page_size).
		From((page-1)*page_size).
		Sort("modify_time", false).
		Pretty(true)
	if cate_id > 0 {
		q := elastic.NewTermQuery("cate_id", cate_id)
		query = query.Query(q)
	}
	result, err := query.Do(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var typ model.Articles
	for _, item := range result.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		t := item.(model.Articles)
		t.ModifyTime = helper.DateToDateTime(t.ModifyTime)
		t.CreateTime = helper.DateToDateTime(t.CreateTime)
		articles = append(articles, t)

	}
	return articles, nil
}

//获取文章详情
func (Articles) GetArticleDetail(article_id int) *ArticleDetails {
	article_content := model.ArticlesContents{}
	article_details := ArticleDetails{}
	db := mysql.Default().GetConn()
	defer db.Close()
	db.Where("article_id = ?", article_id).First(&article_details)
	db.Where("article_id = ?", article_id).First(&article_content)
	article_details.Contents = article_content.Contents
	article_details.ShowType = article_content.ShowType
	return &article_details
}

//获取文章类型
func (Articles) GetArticleCate() []model.ArticlesCate {
	article_cates := make([]model.ArticlesCate, 0)
	db := mysql.Default().GetConn()
	defer db.Close()
	db.Order("orderby asc").Find(&article_cates)
	return article_cates
}
