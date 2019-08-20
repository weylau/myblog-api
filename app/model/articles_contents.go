package model

type ArticlesContents struct {
	ArticleId int32  `json:"article_id"`
	Contents  string `json:"contents"`
}

func (ArticlesContents) TableName() string {
	return "mb_articles_contents"
}
