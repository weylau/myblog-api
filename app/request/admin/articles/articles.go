package articles

type AddRequest struct {
	Title       string `json:"title" validate:"gt=4"`
	CateId      int    `json:"cate_id" validate:"gt=0"`
	Description string `json:"description"`
	Keywords    string `json:"keywords"`
	Contents    string `json:"contents" validate:"required"`
	ImgPath     string `json:"img_path"`
	PublishTime string `json:"publish_time" validate:"required"`
	ShowType    int    `json:"show_type" validate:"required"`
	Status      int    `json:"status" validate:"required"`
	OpId        int
	OpUser      string
}

type UpdateRequest struct {
	Title       string `json:"title" validate:"gt=4"`
	CateId      int    `json:"cate_id" validate:"gt=0"`
	Description string `json:"description"`
	Keywords    string `json:"keywords"`
	Contents    string `json:"contents" validate:"required"`
	ImgPath     string `json:"img_path"`
	PublishTime string `json:"publish_time" validate:"required"`
	ShowType    int    `json:"show_type" validate:"required"`
	Status      int    `json:"status" validate:"required"`
	OpId        int
	OpUser      string
}
