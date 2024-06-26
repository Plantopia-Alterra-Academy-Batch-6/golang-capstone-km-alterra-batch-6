package article

func RequestTOEntityArticle(article *RequestArticle) *Article {
	return &Article{
		Title:   article.Title,
		Content: article.Content,
		Image:   article.Image,
	}
}

func ArticleTOResponse(article *Article) *ResponseArticle {
	return &ResponseArticle{
		ID:        article.ID,
		Title:     article.Title,
		Content:   article.Content,
		Image:     article.Image,
		CreatedAt: article.CreatedAt,
	}
}
