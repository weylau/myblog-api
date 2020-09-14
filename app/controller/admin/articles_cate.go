package admin

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/juju/errors"
	"myblog-api/app/loger"
	"myblog-api/app/protocol"
	"myblog-api/app/request/admin/articles_cate"
	"myblog-api/app/service/admin"
	"myblog-api/app/validate"
	"net/http"
	"strconv"
)
type ArticlesCate struct {
}

func (this *ArticlesCate) DeleteCache(c *gin.Context) {
	resp := &protocol.Resp{Ret: 0, Msg: "", Data: ""}
	serv := admin.ArticlesCate{}
	if(!serv.DeleteCache()) {
		resp.Ret = -1
		resp.Msg = "删除失败"
		c.JSON(http.StatusOK, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}

//文章类型
func (this *ArticlesCate) GetList(c *gin.Context) {
	resp := &protocol.Resp{Ret: 0, Msg: "", Data: ""}
	serv := admin.ArticlesCate{}
	resp = serv.GetArticleCate()
	c.JSON(http.StatusOK, resp)
}


func (this *ArticlesCate) Add(c *gin.Context) {
	resp := &protocol.Resp{Ret: 0, Msg: "", Data: ""}
	var addRequest articles_cate.AddRequest
	err := c.ShouldBindJSON(&addRequest)
	jsonstr, _ := json.Marshal(addRequest)
	loger.Loger.Info("ArticlesCate-Add-Params:", string(jsonstr))
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

	serv := admin.ArticlesCate{}
	resp = serv.Add(&addRequest)
	c.JSON(http.StatusOK, resp)
}

func (this *ArticlesCate) Update(c *gin.Context) {
	resp := &protocol.Resp{Ret: 0, Msg: "", Data: ""}
	cate_id, err := strconv.Atoi(c.Param("cate_id"))
	if err != nil {
		loger.Loger.Info(errors.ErrorStack(errors.Trace(err)))
		resp.Ret = -1
		resp.Msg = "参数错误"
		c.JSON(http.StatusOK, resp)
		return
	}
	var updateRequest articles_cate.UpdateRequest
	err = c.ShouldBindJSON(&updateRequest)
	jsonstr, _ := json.Marshal(updateRequest)
	loger.Loger.Info("ArticlesCate-Update-Params:", string(jsonstr))
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

	serv := admin.ArticlesCate{}
	resp = serv.Update(cate_id, &updateRequest)
	c.JSON(http.StatusOK, resp)
}

func (this *ArticlesCate) Delete(c *gin.Context) {
	resp := &protocol.Resp{Ret: 0, Msg: "", Data: ""}
	cate_id, err := strconv.Atoi(c.Param("cate_id"))
	if err != nil {
		loger.Loger.Info(errors.ErrorStack(errors.Trace(err)))
		resp.Ret = -1
		resp.Msg = "参数错误"
		c.JSON(http.StatusOK, resp)
		return
	}
	serv := admin.ArticlesCate{}
	resp = serv.Delete(cate_id)
	c.JSON(http.StatusOK, resp)
}