package front

import (
	"github.com/gin-gonic/gin"
	"github.com/weylau/myblog-api/app/protocol"
	"github.com/weylau/myblog-api/app/service/front"
	"net/http"
	"strconv"
)

type Articles struct {
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
	article_serv := front.Articles{}
	article_list, err := article_serv.GetListForEs(page, page_size, cate_id, []string{"article_id", "cate_id", "title", "description", "modify_time"})
	if err != nil {
		resp.Ret = -1
		resp.Msg = "系统错误"
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = article_list
	c.JSON(http.StatusOK, resp)
}

//文章详情
func (Articles) GetDetail(c *gin.Context) {
	resp := protocol.Resp{Ret: 0, Msg: "", Data: ""}
	id, err := strconv.Atoi(c.DefaultQuery("id", "0"))

	article_serv := front.Articles{}
	if err != nil || id <= 0 {
		resp.Ret = -1
		resp.Msg = "参数错误"
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = article_serv.GetArticleDetail(id)

	c.JSON(http.StatusOK, resp)
}

//文章类型
func (Articles) GetCate(c *gin.Context) {
	resp := protocol.Resp{Ret: 0, Msg: "", Data: ""}
	article_serv := front.Articles{}
	resp.Data = article_serv.GetArticleCate()
	c.JSON(http.StatusOK, resp)
}
