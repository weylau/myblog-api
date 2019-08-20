package front

import (
	"github.com/gin-gonic/gin"
	"github.com/weylau/myblog-api/app/controllers"
	"github.com/weylau/myblog-api/app/service"
	"net/http"
	"strconv"
)

//文章列表
func GetList(c *gin.Context) {
	resp := controllers.Resp{Ret: 0, Msg: "", Data: ""}
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		page = 1
	}
	page_size, err := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if err != nil {
		page_size = 10
	}
	article_list, err := service.GetList(page, page_size, []string{"article_id", "cate_id", "title", "description", "modify_time"})
	if err != nil {
		resp.Ret = -1
		resp.Msg = "系统错误"
		goto END
	}
	resp.Data = article_list
END:
	c.JSON(http.StatusOK, resp)
}

//文章详情
func GetDetail(c *gin.Context) {
	resp := controllers.Resp{Ret: 0, Msg: "", Data: ""}
	id, err := strconv.Atoi(c.DefaultQuery("id", "0"))
	if err != nil || id <= 0 {
		resp.Ret = -1
		resp.Msg = "参数错误"
		goto END
	}
	resp.Data = service.GetArticleDetail(id)
END:
	c.JSON(http.StatusOK, resp)
}
