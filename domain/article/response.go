package article

// CreateArticleRequest is model for creating article.
type CreateArticleResponse struct {
	Ok    bool `json:"ok"`
	Message string `json:"message"`
}

// EditArticleRequest is model for modified article.
type EditArticleResponse struct {
	Ok    bool `json:"ok"`
	Message string `json:"message"`
}