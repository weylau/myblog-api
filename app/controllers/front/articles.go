package front

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/weylau/myblog-api/app/controllers"
	"github.com/weylau/myblog-api/app/model"
	"net/http"
	"strconv"
)

//文章列表
func GetList(c *gin.Context) {
	resp := controllers.Resp{Ret: 0, Msg: ""}
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		page = 1
	}
	page_size, err := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if err != nil {
		page_size = 10
	}
	article_model := model.Articles{}
	article_list, err := article_model.GetList(page, page_size)
	if err != nil {
		fmt.Println(err)
	}
	resp.Data = article_list
	c.JSON(http.StatusOK, resp)
}

//文章详情
func GetDetail(c *gin.Context) {
	id := c.Query("id")
	ret := fmt.Sprintf("%s:%s", "GetDetail", id)
	c.String(http.StatusOK, ret)
}
