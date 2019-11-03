package admin

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/weylau/myblog-api/app/helpers"
	"github.com/weylau/myblog-api/app/protocol"
	"github.com/weylau/myblog-api/app/service/admin"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Articles struct {
}

type AddParams struct {
	Title       string `json:"title"`
	CateId      int    `json:"cate_id"`
	Descreption string `json:"descreption"`
	Keywords    string `json:"keywords"`
	Contents    string `json:"contents"`
	ImgPath     string `json:"img_path"`
	PublishTime string `json:"publish_time"`
	ShowType    int    `json:"show_type"`
}

//文章列表
func (Articles) Add(c *gin.Context) {
	resp := &protocol.Resp{Ret: -1, Msg: "", Data: ""}
	helper := helpers.Helpers{}
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		resp.Ret = -1
		resp.Msg = "参数错误1"
		c.JSON(http.StatusOK, resp)
		return
	}
	addParams := &AddParams{}
	err = json.Unmarshal(data, addParams)
	if err != nil {
		resp.Ret = -1
		resp.Msg = "参数错误2" + err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	if addParams.CateId <= 0 {
		resp.Ret = -1
		resp.Msg = "CateId参数错误"
		c.JSON(http.StatusOK, resp)
		return
	}
	if addParams.Title == "" || addParams.Contents == "" {
		resp.Ret = -1
		resp.Msg = "参数错误"
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
	params.Description = addParams.Descreption
	params.Keywords = addParams.Keywords
	params.Contents = addParams.Contents
	params.ImgPath = addParams.ImgPath
	params.PublishTime = addParams.PublishTime
	params.ShowType = addParams.ShowType
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
func (Articles) GetList(c *gin.Context) {
	resp := protocol.Resp{Ret: 0, Msg: "", Data: ""}
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
	article_list, err := article_serv.GetList(page, page_size, cate_id, []string{"article_id", "cate_id", "title", "description", "op_user", "modify_time"})
	if err != nil {
		resp.Ret = -1
		resp.Msg = "系统错误"
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = article_list
	c.JSON(http.StatusOK, resp)
}
