package model

import (
	"github.com/weylau/myblog-api/app/db"
)

type Articles struct {
	ArticleId   int32  `db:"article_id" json:"article_id"`
	CateId      int32  `db:"cate_id" json:"cate_id"`
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
	Keywords    string `db:"keywords" json:"keywords"`
	ImgPath     string `db:"img_path" json:"img_path"`
	ModifyTime  string `db:"modify_time" json:"modify_time"`
	OpId        int32  `db:"op_id" json:"op_id"`
	OpUser      string `db:"op_user" json:"op_user"`
	CreateTime  string `db:"create_time" json:"create_time"`
}

func (this *Articles) GetList(page int, page_size int) ([]*Articles, error) {
	db := db.DBConn()
	defer db.Close()
	offset := (page - 1) * page_size
	rows, err := db.Query("select article_id,cate_id,title from mb_articles limit ?,?", offset, page_size)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	var listArticles []*Articles
	for rows.Next() {
		article := &Articles{}

		err = rows.Scan(
			&article.ArticleId,
			&article.CateId,
			&article.Title,
		//&article.Description,
		//&article.Keywords,
		//&article.ImgPath,
		//&article.ModifyTime,
		//&article.OpId,
		//&article.OpUser,
		//&article.CreateTime,
		)
		if err != nil {
			panic(err.Error())
		}

		listArticles = append(listArticles, article)
	}
	return listArticles, nil
}
