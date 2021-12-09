package article

// CreateArticleRequest is model for creating article.
type CreateArticleRequest struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Content  string `json:"content"`
	Status ArticleStatus `json:"status"`
	Email string `json:"email"`
}

// EditArticleRequest is model for modified article.
type EditArticleRequest struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Content  string `json:"content"`
}
