package admin

import (
	"github.com/juju/errors"
	"myblog-api/app/config"
	"myblog-api/app/db/mysql"
	"myblog-api/app/db/redis"
	"myblog-api/app/loger"
	"myblog-api/app/model"
	"myblog-api/app/protocol"
	"myblog-api/app/request/admin/articles_cate"
)

type CateParams struct {
	Name        string `json:"name"`
	CName       string `json:"c_name"`
	Orderby     int    `json:"orderby"`
}

type ArticlesCate struct {
}

//获取文章类型
func (this *ArticlesCate) GetArticleCate() (resp *protocol.Resp) {
	resp = &protocol.Resp{Ret: -1, Msg: "", Data: ""}

	article_cates, err := this.articleCate();
	if err != nil {
		resp.Msg = "系统错误"
		return resp
	}

	resp.Ret = 0
	resp.Data = article_cates
	return resp
}

//添加文章分类
func (this *ArticlesCate) Add(cateParams *articles_cate.AddRequest)(resp *protocol.Resp) {
	resp = &protocol.Resp{Ret: 0, Msg: "", Data: ""}
	if !this.DeleteCache() {
		resp.Msg = "添加失败，请重试！"
		return resp
	}
	db := mysql.MysqlDB.GetConn()
	cate := model.ArticlesCate{}
	cate.CName = cateParams.CName
	cate.Name = cateParams.Name
	cate.Orderby = cateParams.Orderby
	err := db.Model(model.Articles{}).Create(&cate).Error
	if err != nil {
		loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
		resp.Msg = "系统错误"
		resp.Ret = -1
	}
	return resp
}

//编辑文章分类
func (this *ArticlesCate) Update(cateId int, cateParams *articles_cate.UpdateRequest)(resp *protocol.Resp) {
	resp = &protocol.Resp{Ret: 0, Msg: "", Data: ""}
	if !this.DeleteCache() {
		resp.Msg = "更新失败，请重试！"
		return resp
	}
	db := mysql.MysqlDB.GetConn()
	count := 0
	if err := db.Model(model.ArticlesCate{}).Where("cate_id = ?", cateId).Count(&count).Error; err != nil {
		loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
		resp.Msg = "系统错误"
		return resp
	}

	if count <= 0 {
		resp.Msg = "分类不存在"
		return resp
	}
	cate := model.ArticlesCate{}
	cate.CName = cateParams.CName
	cate.Name = cateParams.Name
	cate.Orderby = cateParams.Orderby
	err := db.Model(model.ArticlesCate{}).Where("cate_id = ?", cateId).Update(&cate).Error
	if err != nil {
		loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
		resp.Msg = "系统错误"
		resp.Ret = -1
	}
	return resp
}

func (this *ArticlesCate) articleCate() ([]model.ArticlesCate, error) {
	article_cates := make([]model.ArticlesCate, 0)
	db := mysql.MysqlDB.GetConn()
	if err := db.Order("orderby asc").Find(&article_cates).Error; err != nil {
		loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
		return nil, err
	}
	return article_cates, nil
}

//清空cate缓存
func (this *ArticlesCate) DeleteCache() (bool){
	cacheKey := "article_cates_"+config.Configs.RedisCacheVersion
	redisConn := redis.RedisClient.Pool.Get()
	if err := redisConn.Err(); err != nil {
		loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
		return false
	}
	_, err := redisConn.Do("del",cacheKey)
	if  err != nil {
		loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
		return false
	}
	return true;
}


func (this *ArticlesCate) Delete(id int) (resp *protocol.Resp) {
	resp = &protocol.Resp{Ret: -1, Msg: "", Data: ""}
	if !this.DeleteCache() {
		resp.Msg = "删除失败，请重试！"
		return resp
	}
	db := mysql.MysqlDB.GetConn()
	if err := db.Where("cate_id = ?", id).Delete(&model.ArticlesCate{}).Error; err != nil {
		loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
		resp.Msg = "系统错误"
		return resp
	}
	resp.Ret = 0
	return resp
}