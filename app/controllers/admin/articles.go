package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/weylau/myblog-api/app/protocol"
	"github.com/weylau/myblog-api/app/service/admin"
	"net/http"
	"strconv"
)

type Articles struct {
}

//文章列表
func (Articles) Add(c *gin.Context) {
	resp := &protocol.Resp{Ret: -1, Msg: "", Data: ""}
	params := &admin.ArticleParams{}
	cate_id, err := strconv.Atoi(c.DefaultPostForm("cate_id", "0"))
	if err != nil {
		resp.Msg = "系统错误:" + err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	params.CateId = cate_id

	params.Title = c.DefaultPostForm("title", "")
	params.Description = c.DefaultPostForm("description", "")
	params.Keywords = c.DefaultPostForm("keywords", "")
	params.Contents = c.DefaultPostForm("contents", "")
	params.ImgPath = c.DefaultPostForm("img_path", "")
	show_type, err := strconv.Atoi(c.DefaultPostForm("show_type", "0"))
	if err != nil {
		resp.Msg = "系统错误:" + err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	params.ShowType = show_type
	if err != nil {
		resp.Msg = "系统错误:" + err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	article_serv := admin.Articles{}
	resp = article_serv.Add(params)
	c.JSON(http.StatusOK, resp)
}
