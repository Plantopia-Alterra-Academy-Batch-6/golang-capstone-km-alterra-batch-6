package article

type RequestArticle struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
	Image   string `json:"image" validate:"required, image"`
}
