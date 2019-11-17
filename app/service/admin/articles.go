package admin

import (
	"myblog-api/app/db/mysql"
	"myblog-api/app/model"
	"myblog-api/app/protocol"
	"time"
)

type ArticleParams struct {
	CateId      int    `json:"cate_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Keywords    string `json:"keywords"`
	ImgPath     string `json:"img_path"`
	OpId        int    `json:"op_id"`
	OpUser      string `json:"op_user"`
	Contents    string `json:"contents"`
	ShowType    int    `json:"show_type"`
	PublishTime string `json:"publish_time"`
}

type Detail struct {
	Id          int    `json:"id"`
	CateId      int    `json:"cate_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Keywords    string `json:"keywords"`
	ImgPath     string `json:"img_path"`
	Contents    string `json:"contents"`
	ShowType    int    `json:"show_type"`
	PublishTime string `json:"publish_time"`
}

type Articles struct {
}

type ArticleList struct {
	Total    int              `json:"total"`
	Datalist []model.Articles `json:"datalist"`
}

//添加文章
func (Articles) Add(params *ArticleParams) (resp *protocol.Resp) {
	resp = &protocol.Resp{Ret: -1, Msg: "", Data: ""}
	articles := model.Articles{
		CateId:      params.CateId,
		Title:       params.Title,
		Description: params.Description,
		Keywords:    params.Keywords,
		ImgPath:     params.ImgPath,
		OpId:        params.OpId,
		OpUser:      params.OpUser,
		ModifyTime:  params.PublishTime,
		CreateTime:  time.Now().Format("2006-01-02 15:04:05"),
	}

	articles_contents := model.ArticlesContents{
		ShowType: params.ShowType,
		Contents: params.Contents,
	}
	//if articles_contents.GetShowTypeName() == "" {
	//	resp.Msg = "文章内容显示类型错误"
	//	return resp
	//}
	//添加articles_contents
	db := mysql.Default().GetConn()
	defer db.Close()
	// 开始事务
	tx := db.Begin()
	//添加articles
	err := db.Model(model.Articles{}).Create(&articles).Error
	if err != nil {
		resp.Msg = "系统错误：" + err.Error()
		tx.Rollback()
		return resp
	}
	//获取插入记录的Id
	var article_id []int
	db.Raw("select LAST_INSERT_ID() as id").Pluck("article_id", &article_id)
	articles_contents.ArticleId = article_id[0]
	err = db.Create(&articles_contents).Error
	if err != nil {
		resp.Msg = "系统错误：" + err.Error()
		tx.Rollback()
		return resp
	}
	//提交事务
	tx.Commit()
	resp.Ret = 0
	return resp
}

//更新文章
func (Articles) Update(id int, params *ArticleParams) (resp *protocol.Resp) {
	resp = &protocol.Resp{Ret: -1, Msg: "", Data: ""}
	//查询ID是否存在
	db := mysql.Default().GetConn()
	defer db.Close()
	count := 0
	if err := db.Model(model.Articles{}).Where("article_id = ?", id).Count(&count).Error; err != nil {
		resp.Msg = "系统错误：" + err.Error()
		return resp
	}

	if count <= 0 {
		resp.Msg = "文章不存在"
		return resp
	}
	articles := model.Articles{
		CateId:      params.CateId,
		Title:       params.Title,
		Description: params.Description,
		Keywords:    params.Keywords,
		ImgPath:     params.ImgPath,
		OpId:        params.OpId,
		OpUser:      params.OpUser,
		ModifyTime:  params.PublishTime,
	}

	articles_contents := model.ArticlesContents{
		ShowType: params.ShowType,
		Contents: params.Contents,
	}
	//if articles_contents.GetShowTypeName() == "" {
	//	resp.Msg = "文章内容显示类型错误"
	//	return resp
	//}
	// 开始事务
	tx := db.Begin()
	//添加articles
	err := db.Model(model.Articles{}).Where("article_id = ?", id).Update(&articles).Error
	if err != nil {
		resp.Msg = "系统错误：" + err.Error()
		tx.Rollback()
		return resp
	}
	//获取插入记录的Id
	err = db.Model(model.ArticlesContents{}).Where("article_id = ?", id).Updates(&articles_contents).Error
	if err != nil {
		resp.Msg = "系统错误：" + err.Error()
		tx.Rollback()
		return resp
	}
	//提交事务
	tx.Commit()
	resp.Ret = 0
	return resp
}

//分页获取文章列表
func (Articles) GetList(page int, page_size int, cate_id int, fields []string) (*ArticleList, error) {
	db := mysql.Default().GetConn()
	defer db.Close()
	offset := (page - 1) * page_size
	article_list := &ArticleList{}
	articles := make([]model.Articles, 0)
	total := 0
	if cate_id > 0 {
		db = db.Where("cate_id = ?", cate_id)
	}
	db.Model(&model.Articles{}).Count(&total)
	db.Select(fields).Offset(offset).Limit(page_size).Order("article_id desc").Find(&articles)
	article_list.Datalist = articles
	article_list.Total = total
	return article_list, nil
}

//删除文章
func (Articles) Delete(id int) (bool, error) {
	db := mysql.Default().GetConn()
	defer db.Close()
	if err := db.Where("article_id = ?", id).Delete(&model.Articles{}).Error; err != nil {
		return false, err
	}
	return true, nil
}

//文章详情
func (Articles) Detail(id int) (*Detail, error) {
	db := mysql.Default().GetConn()
	defer db.Close()
	article := &model.Articles{}
	article_content := &model.ArticlesContents{}

	if err := db.Where("article_id = ?", id).Find(article).Error; err != nil {
		return nil, err
	}
	if err := db.Where("article_id = ?", id).Find(article_content).Error; err != nil {
		return nil, err
	}
	detail := &Detail{}
	detail.Title = article.Title
	detail.Id = id
	detail.CateId = article.CateId
	detail.Description = article.Description
	detail.Keywords = article.Keywords
	detail.ImgPath = article.ImgPath
	detail.PublishTime = article.ModifyTime
	detail.Contents = article_content.Contents
	detail.ShowType = article_content.ShowType
	return detail, nil
}
