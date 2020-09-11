package front

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic"
	"myblog-api/app/config"
	"myblog-api/app/db/es"
	"myblog-api/app/db/mysql"
	"myblog-api/app/db/redis"
	"myblog-api/app/helper"
	"myblog-api/app/loger"
	"myblog-api/app/model"
	"myblog-api/app/protocol"
	"reflect"
	red "github.com/gomodule/redigo/redis"
	"github.com/juju/errors"
	"strconv"

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
func (this *Articles) GetListForMysql(page int, page_size int, cate_id int, fields []string) (resp *protocol.Resp) {
	resp = &protocol.Resp{Ret: -1, Msg: "", Data: ""}
	db := mysql.MysqlDB.GetConn()
	offset := (page - 1) * page_size
	articles := make([]model.Articles, 0)
	db = db.Where("status = ?", 1)
	if cate_id > 0 {
		db = db.Where("cate_id = ?", cate_id)
	}
	if err := db.Select(fields).Offset(offset).Limit(page_size).Order("article_id desc").Find(&articles).Error; err != nil {
		loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
		resp.Msg = "系统错误"
		return resp
	}
	resp.Ret = 0
	resp.Data = articles
	return resp
}

//分页获取文章列表es
func (this *Articles) GetListForEs(page int, page_size int, cate_id int, fields []string) (resp *protocol.Resp) {
	resp = &protocol.Resp{Ret: -1, Msg: "", Data: ""}
	esclient,err := es.Default()
	if err != nil {
		loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
		resp.Ret = -999
		resp.Msg = "系统错误"
		return resp
	}
	esconn := esclient.GetConn()
	ctx := context.Background()
	articles := make([]model.Articles, 0)

	query := esconn.Search().
		Index("myblog").
		Type("mb_articles").
		Size(page_size).
		From((page-1)*page_size).
		Sort("modify_time", false).
		Pretty(true)
	boolQuery := elastic.NewBoolQuery()
	searchQuery := boolQuery.Must(elastic.NewTermQuery("status", 1))
	if cate_id > 0 {
		searchQuery = searchQuery.Filter(elastic.NewTermQuery("cate_id", cate_id))
	}
	query = query.Query(searchQuery)
	result, err := query.Do(ctx)
	if err != nil {
		loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
		resp.Ret = -999
		resp.Msg = "系统错误"
		return resp
	}
	var typ model.Articles
	for _, item := range result.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		t := item.(model.Articles)
		t.ModifyTime = helper.DateToDateTime(t.ModifyTime)
		t.CreateTime = helper.DateToDateTime(t.CreateTime)
		articles = append(articles, t)

	}
	resp.Ret = 0
	resp.Data = articles
	return resp
}

//获取文章详情
func (this *Articles) GetArticleDetail(article_id int) (resp *protocol.Resp) {
	resp = &protocol.Resp{Ret: -1, Msg: "", Data: ""}
	article_content := model.ArticlesContents{}

	redisConn := redis.RedisClient.Pool.Get()
	cacheKey := "article_"+config.Configs.RedisCacheVersion+":"+strconv.Itoa(article_id)
	if err := redisConn.Err(); err != nil {
		loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
	} else {
		cateCache, err := red.String(redisConn.Do("GET", cacheKey))
		if err != nil {
			loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
		} else {
			if cateCache != "" {
				article_details := ArticleDetails{}
				err := json.Unmarshal([]byte(cateCache), &article_details)
				if err != nil {
					loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
				}
				resp.Ret = 0
				resp.Data = article_details
				return resp
			}
		}
	}
	article_details := ArticleDetails{}
	db := mysql.MysqlDB.GetConn()
	if err := db.Where("article_id = ?", article_id).Where("status = ?", 1).First(&article_details).Error; err != nil {
		loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
		resp.Msg = "文章不存在"
		return resp
	}
	if article_details.ArticleId <= 0 {
		resp.Msg = "文章不存在"
		return resp
	}
	if err := db.Where("article_id = ?", article_id).First(&article_content).Error; err != nil {
		loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
		resp.Msg = "系统错误"
		return resp
	}
	resp.Ret = 0
	article_details.Contents = article_content.Contents
	article_details.ShowType = article_content.ShowType
	cacheData,err := json.Marshal(article_details)
	if err == nil {
		redisConn.Do("set",cacheKey, string(cacheData))
		redisConn.Do("expire",cacheKey, 86400*3)
	}
	resp.Data = article_details
	return resp
}

//获取文章类型
func (this *Articles) GetArticleCate() (resp *protocol.Resp) {
	resp = &protocol.Resp{Ret: -1, Msg: "", Data: ""}
	redisConn := redis.RedisClient.Pool.Get()
	cacheKey := "article_cates_"+config.Configs.RedisCacheVersion
	if err := redisConn.Err(); err != nil {
		loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
	} else {
		cateCache, err := red.String(redisConn.Do("GET", cacheKey))
		if err != nil {
			loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
		} else {
			if cateCache != "" {
				article_cates := make([]model.ArticlesCate, 0)
				err := json.Unmarshal([]byte(cateCache), &article_cates)
				if err != nil {
					loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
				}
				resp.Ret = 0
				resp.Data = article_cates
				return resp
			}
		}
	}

	article_cates, err := this.articleCate();
	if  err != nil {
		resp.Msg = "系统错误"
		return resp
	}

	if len(article_cates) > 0 {
		cacheData,err := json.Marshal(article_cates)
		if err == nil {
			redisConn.Do("set",cacheKey, string(cacheData))
			redisConn.Do("expire",cacheKey, 86400*3)
		}

	}

	resp.Ret = 0
	resp.Data = article_cates
	return resp
}

func (this *Articles) articleCate() ([]model.ArticlesCate, error){
	article_cates := make([]model.ArticlesCate, 0)
	db := mysql.MysqlDB.GetConn()
	if err := db.Order("orderby asc").Find(&article_cates).Error; err != nil {
		loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
		return nil,err
	}
	return article_cates,nil
}
