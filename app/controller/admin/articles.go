package admin

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"myblog-api/app/helper"
	"myblog-api/app/loger"
	"myblog-api/app/protocol"
	"myblog-api/app/service/admin"
	"myblog-api/app/validate"
	"net/http"
	"strconv"
	"myblog-api/app/config"
	"myblog-api/app/db/redis"
	"github.com/juju/errors"
)

type Articles struct {
}


type AddParams struct {
	Title       string `json:"title" validate:"gt=4"`
	CateId      int    `json:"cate_id" validate:"gt=0"`
	Description string `json:"description"`
	Keywords    string `json:"keywords"`
	Contents    string `json:"contents" validate:"required"`
	ImgPath     string `json:"img_path"`
	PublishTime string `json:"publish_time" validate:"required"`
	ShowType    int    `json:"show_type" validate:"required"`
	Status      int    `json:"status" validate:"required"`
}

//添加文章
func (this *Articles) Add(c *gin.Context) {
	resp := &protocol.Resp{Ret: -1, Msg: "", Data: ""}
	var addParams AddParams
	err := c.ShouldBindJSON(&addParams)
	jsonstr, _ := json.Marshal(addParams)
	loger.Loger.Info("Articles-Add-Params:", string(jsonstr))
	if err != nil {
		loger.Loger.Info(errors.ErrorStack(errors.Trace(err)))
		resp.Ret = -1
		resp.Msg = "参数错误"
		c.JSON(http.StatusOK, resp)
		return
	}

	validator, _ := validate.Default()
	if check := validator.CheckStruct(addParams); !check {
		resp.Ret = -1
		resp.Msg = validator.GetOneError()
		c.JSON(http.StatusOK, resp)
		return
	}

	if !helper.IsTimeStr(addParams.PublishTime) {
		resp.Ret = -1
		resp.Msg = "发布时间格式错误"
		c.JSON(http.StatusOK, resp)
		return
	}
	if addParams.ShowType <= 0 {
		addParams.ShowType = 2
	}
	params := &admin.ArticleParams{}
	params.CateId = addParams.CateId
	params.Title = addParams.Title
	params.Description = addParams.Description
	params.Keywords = addParams.Keywords
	params.Contents = addParams.Contents
	params.ImgPath = addParams.ImgPath
	params.PublishTime = addParams.PublishTime
	params.ShowType = addParams.ShowType
	params.Status = addParams.Status
	admin_id, _ := c.Get("admin_id")
	username, _ := c.Get("username")
	aop_id, _ := strconv.Atoi(admin_id.(string))
	params.OpId = aop_id
	params.OpUser = username.(string)

	article_serv := admin.Articles{}
	resp = article_serv.Add(params)
	c.JSON(http.StatusOK, resp)
}

//文章列表
func (this *Articles) GetList(c *gin.Context) {
	resp := &protocol.Resp{Ret: 0, Msg: "2", Data: ""}
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		page = 1
	}
	page_size, err := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if err != nil {
		page_size = 10
	}
	cate_id, err := strconv.Atoi(c.DefaultQuery("cate_id", "0"))
	if err != nil {
		cate_id = 0
	}
	article_serv := admin.Articles{}
	resp = article_serv.GetList(page, page_size, cate_id, []string{"article_id", "cate_id", "title", "description", "op_user", "modify_time", "status"})
	c.JSON(http.StatusOK, resp)
}

//删除文章
func (this *Articles) Delete(c *gin.Context) {
	resp := &protocol.Resp{Ret: 0, Msg: "", Data: ""}
	article_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		loger.Loger.Info(errors.ErrorStack(errors.Trace(err)))
		resp.Ret = -1
		resp.Msg = "参数错误"
		c.JSON(http.StatusOK, resp)
		return
	}
	article_serv := admin.Articles{}
	resp = article_serv.Delete(article_id)
	c.JSON(http.StatusOK, resp)
}

//更新文章
func (this *Articles) Update(c *gin.Context) {
	resp := &protocol.Resp{Ret: -1, Msg: "", Data: ""}
	article_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		loger.Loger.Info(errors.ErrorStack(errors.Trace(err)))
		resp.Ret = -1
		resp.Msg = "参数错误"
		c.JSON(http.StatusOK, resp)
		return
	}
	var addParams AddParams
	err = c.ShouldBindJSON(&addParams)
	jsonstr, _ := json.Marshal(addParams)
	loger.Loger.Info("Articles-Update-Params:", string(jsonstr))
	if err != nil {
		loger.Loger.Info(errors.ErrorStack(errors.Trace(err)))
		resp.Ret = -1
		resp.Msg = "参数错误"
		c.JSON(http.StatusOK, resp)
		return
	}

	validator, _ := validate.Default()
	if check := validator.CheckStruct(addParams); !check {
		resp.Ret = -1
		resp.Msg = validator.GetOneError()
		c.JSON(http.StatusOK, resp)
		return
	}

	if !helper.IsTimeStr(addParams.PublishTime) {
		resp.Ret = -1
		resp.Msg = "发布时间格式错误"
		c.JSON(http.StatusOK, resp)
		return
	}
	if addParams.ShowType <= 0 {
		addParams.ShowType = 2
	}
	params := &admin.ArticleParams{}
	params.CateId = addParams.CateId
	params.Title = addParams.Title
	params.Description = addParams.Description
	params.Keywords = addParams.Keywords
	params.Contents = addParams.Contents
	params.ImgPath = addParams.ImgPath
	params.PublishTime = addParams.PublishTime
	params.ShowType = addParams.ShowType
	params.Status = addParams.Status
	admin_id, _ := c.Get("admin_id")
	username, _ := c.Get("username")
	aop_id, _ := strconv.Atoi(admin_id.(string))
	params.OpId = aop_id
	params.OpUser = username.(string)

	article_serv := admin.Articles{}
	resp = article_serv.Update(article_id, params)
	c.JSON(http.StatusOK, resp)
}

func (this *Articles) Show(c *gin.Context) {
	resp := &protocol.Resp{Ret: 0, Msg: "", Data: ""}
	article_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
		resp.Ret = -1
		resp.Msg = "参数错误"
		c.JSON(http.StatusOK, resp)
		return
	}
	article_serv := admin.Articles{}
	resp = article_serv.Detail(article_id)
	c.JSON(http.StatusOK, resp)
}

func (this *Articles) DeleteCache(c *gin.Context) {
	resp := &protocol.Resp{Ret: 0, Msg: "", Data: ""}
	cacheKey := "article_cates_"+config.Configs.RedisCacheVersion
	redisConn := redis.RedisClient.Pool.Get()
	if err := redisConn.Err(); err != nil {
		loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
		resp.Ret = -1
		resp.Msg = "系统错误"
		return
	}
	_, err := redisConn.Do("del",cacheKey)
	if  err != nil {
		loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
		resp.Ret = -1
		resp.Msg = "系统错误"
		return
	}
	c.JSON(http.StatusOK, resp)
}
