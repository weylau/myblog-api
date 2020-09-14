package articles_cate

type AddRequest struct {
	Name        string `json:"name" validate:"required"`
	CName       string `json:"c_name" validate:"required"`
	Orderby     int    `json:"orderby" validate:"required"`
}

type UpdateRequest struct {
	Name        string `json:"name" validate:"required"`
	CName       string `json:"c_name" validate:"required"`
	Orderby     int    `json:"orderby" validate:"required"`
}