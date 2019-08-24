package model

type ArticlesContents struct {
	ArticleId int    `json:"article_id"`
	ShowType  int    `json:"show_type"`
	Contents  string `json:"contents"`
}

func (ArticlesContents) TableName() string {
	return "mb_articles_contents"
}

//获取文章显示类型
func (this ArticlesContents) GetShowTypeName() string {
	show_type_name := ""
	switch this.ShowType {
	case 1:
		show_type_name = "html"
	case 2:
		show_type_name = "markdown"
	default:
		show_type_name = ""
	}
	return show_type_name
}
