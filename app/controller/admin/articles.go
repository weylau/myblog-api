package admin

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/juju/errors"
	"myblog-api/app/helper"
	"myblog-api/app/loger"
	"myblog-api/app/protocol"
	"myblog-api/app/request/admin/articles"
	"myblog-api/app/service/admin"
	"myblog-api/app/validate"
	"net/http"
	"strconv"
)

type Articles struct {
}


//添加文章
func (this *Articles) Add(c *gin.Context) {
	resp := &protocol.Resp{Ret: -1, Msg: "", Data: ""}
	var addRequest articles.AddRequest
	err := c.ShouldBindJSON(&addRequest)
	jsonstr, _ := json.Marshal(addRequest)
	loger.Loger.Info("Articles-Add-Params:", string(jsonstr))
	if err != nil {
		loger.Loger.Info(errors.ErrorStack(errors.Trace(err)))
		resp.Ret = -1
		resp.Msg = "参数错误"
		c.JSON(http.StatusOK, resp)
		return
	}

	validator, _ := validate.Default()
	if check := validator.CheckStruct(addRequest); !check {
		resp.Ret = -1
		resp.Msg = validator.GetOneError()
		c.JSON(http.StatusOK, resp)
		return
	}
	admin_id, _ := c.Get("admin_id")
	username, _ := c.Get("username")
	aop_id, _ := strconv.Atoi(admin_id.(string))
	addRequest.OpId = aop_id
	addRequest.OpUser = username.(string)

	if !helper.IsTimeStr(addRequest.PublishTime) {
		resp.Ret = -1
		resp.Msg = "发布时间格式错误"
		c.JSON(http.StatusOK, resp)
		return
	}
	if addRequest.ShowType <= 0 {
		addRequest.ShowType = 2
	}
	serv := admin.Articles{}
	resp = serv.Add(&addRequest)
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
	serv := admin.Articles{}
	resp = serv.GetList(page, page_size, cate_id, []string{"article_id", "cate_id", "title", "description", "op_user", "modify_time", "status"})
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
	serv := admin.Articles{}
	resp = serv.Delete(article_id)
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
	var updateRequest articles.UpdateRequest
	err = c.ShouldBindJSON(&updateRequest)
	jsonstr, _ := json.Marshal(updateRequest)
	loger.Loger.Info("Articles-Update-Params:", string(jsonstr))
	if err != nil {
		loger.Loger.Info(errors.ErrorStack(errors.Trace(err)))
		resp.Ret = -1
		resp.Msg = "参数错误"
		c.JSON(http.StatusOK, resp)
		return
	}

	validator, _ := validate.Default()
	if check := validator.CheckStruct(updateRequest); !check {
		resp.Ret = -1
		resp.Msg = validator.GetOneError()
		c.JSON(http.StatusOK, resp)
		return
	}
	admin_id, _ := c.Get("admin_id")
	username, _ := c.Get("username")
	aop_id, _ := strconv.Atoi(admin_id.(string))
	updateRequest.OpId = aop_id
	updateRequest.OpUser = username.(string)
	if !helper.IsTimeStr(updateRequest.PublishTime) {
		resp.Ret = -1
		resp.Msg = "发布时间格式错误"
		c.JSON(http.StatusOK, resp)
		return
	}
	if updateRequest.ShowType <= 0 {
		updateRequest.ShowType = 2
	}

	serv := admin.Articles{}
	resp = serv.Update(article_id, &updateRequest)
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
	serv := admin.Articles{}
	resp = serv.Detail(article_id)
	c.JSON(http.StatusOK, resp)
}

